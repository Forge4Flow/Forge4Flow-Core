package flow

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	nonce "github.com/forge4flow/forge4flow-core/pkg/authn/nonce"
	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func ApiKeyAndAccountProofAuthMidleware(cfg config.Config, next http.Handler, svcs ...service.Service) (http.Handler, error) {
	var flowService *FlowService
	var nonceService *nonce.NonceService
	for _, svc := range svcs {
		if svc.ID() == service.FlowService {
			flowService = svc.(*FlowService)
		}

		if svc.ID() == service.NonceService {
			nonceService = svc.(*nonce.NonceService)
		}
	}
	forge4FlowConfig, ok := cfg.(config.Forge4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Forge4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authType, auth, err := parseAuthFromRequest(r, []string{service.AuthTypeApiKey})
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		var authInfo *service.AuthInfo
		switch authType {
		case service.AuthTypeApiKey:
			apiKey := auth.(string)
			if !service.SecureCompareEqual(apiKey, forge4FlowConfig.GetAuthentication().ApiKey) {
				service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key"))
				return
			}

			authInfo = &service.AuthInfo{}
		case service.AuthTypeAccountProof:
			accountProof := auth.(AccountProofSpec)
			//TODO: Fix error handling
			validNonce, _ := nonceService.IsValid(r.Context(), accountProof.Nonce)
			if !validNonce {
				service.SendErrorResponse(w, service.NewInvalidRequestError("Invalid Nonce Received"))
				return
			}

			validAccountProof, err := flowService.VerifyAccountProof(r.Context(), accountProof)
			if err != nil {
				service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid Account Proof: %s", err.Error())))
				return
			}

			if validAccountProof {
				authInfo := &service.AuthInfo{
					UserId: accountProof.Address,
				}

				newContext := context.WithValue(r.Context(), service.AuthInfoKey, *authInfo)
				next.ServeHTTP(w, r.WithContext(newContext))
				return
			}

			service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid Account Proof"))
		}

		newContext := context.WithValue(r.Context(), service.AuthInfoKey, *authInfo)
		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}

func parseAuthFromRequest(r *http.Request, validTokenTypes []string) (string, interface{}, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	var authToken interface{}
	if len(authHeaderParts) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	authTokenType := authHeaderParts[0]
	authToken = authHeaderParts[1]

	var isValidTokenType bool
	for _, validTokenType := range validTokenTypes {
		if authTokenType == validTokenType {
			isValidTokenType = true
		}
	}
	if !isValidTokenType {
		var auth AccountProofSpec
		err := service.ParseJSONBody(r.Body, &auth)
		if err != nil {
			return "", "", fmt.Errorf("invalid authorization header or no account proof")
		}

		authTokenType = service.AuthTypeAccountProof
		authToken = auth
	}

	return authTokenType, authToken, nil
}
