package service

import (
	"context"
	"crypto/subtle"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/auth4flow/auth4flow-core/pkg/config"
)

const FirebasePublicKeyUrl = "https://www.googleapis.com/robot/v1/metadata/x509/securetoken@system.gserviceaccount.com"

type key int

const (
	AuthInfoKey key = iota
)

const (
	AuthTypeAccountProof = "AccountProof"
	AuthTypeApiKey       = "ApiKey"
	AuthTypeBearer       = "Bearer"
)

type AuthInfo struct {
	UserId   string
	TenantId string
}

type AuthMiddlewareFunc func(config config.Config, next http.Handler, svcs ...Service) (http.Handler, error)

func ApiKeyAuthMiddleware(cfg config.Config, next http.Handler, svcs ...Service) (http.Handler, error) {
	warrantCfg, ok := cfg.(config.Auth4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Auth4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, tokenString, err := ParseAuthTokenFromRequest(r, []string{AuthTypeApiKey})
		if err != nil {
			SendErrorResponse(w, NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		if !SecureCompareEqual(tokenString, warrantCfg.GetAuthentication().ApiKey) {
			SendErrorResponse(w, NewUnauthorizedError("Invalid API key"))
			return
		}

		newContext := context.WithValue(r.Context(), AuthInfoKey, &AuthInfo{})
		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}

func PassthroughAuthMiddleware(cfg config.Config, next http.Handler, svcs ...Service) (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}), nil
}

// GetAuthInfoFromRequestContext returns the AuthInfo object from the given context
func GetAuthInfoFromRequestContext(context context.Context) *AuthInfo {
	contextVal := context.Value(AuthInfoKey)
	if contextVal != nil {
		authInfo := context.Value(AuthInfoKey).(AuthInfo)
		return &authInfo
	}

	return nil
}

func ParseAuthTokenFromRequest(r *http.Request, validTokenTypes []string) (string, string, error) {
	authHeader := r.Header.Get("Authorization")
	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 {
		return "", "", fmt.Errorf("invalid format")
	}

	authTokenType := authHeaderParts[0]
	authToken := authHeaderParts[1]

	var isValidTokenType bool
	for _, validTokenType := range validTokenTypes {
		if authTokenType == validTokenType {
			isValidTokenType = true
		}
	}
	if !isValidTokenType {
		return "", "", fmt.Errorf("authorization header prefix must be one of: %s", strings.Join(validTokenTypes, ", "))
	}

	return authTokenType, authToken, nil
}

func SecureCompareEqual(given string, actual string) bool {
	if subtle.ConstantTimeEq(int32(len(given)), int32(len(actual))) == 1 {
		return subtle.ConstantTimeCompare([]byte(given), []byte(actual)) == 1
	} else {
		return subtle.ConstantTimeCompare([]byte(actual), []byte(actual)) == 1 && false
	}
}
