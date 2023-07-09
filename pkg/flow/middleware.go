package flow

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	nonce "github.com/auth4flow/auth4flow-core/pkg/authn/nonce"
	"github.com/auth4flow/auth4flow-core/pkg/config"
	"github.com/auth4flow/auth4flow-core/pkg/service"
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
	_, ok := cfg.(config.Auth4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Auth4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var accountProof AccountProofSpec
		err := service.ParseJSONBody(r.Body, &accountProof)
		if err != nil {
			service.SendErrorResponse(w, service.NewInvalidRequestError("Invalid AccountProof Received"))
			return
		}

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
	}), nil
}
