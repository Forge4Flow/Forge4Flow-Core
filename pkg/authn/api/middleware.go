package authn

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func MasterApiKeyAuthMiddleware(cfg config.Config, next http.Handler, svcs ...service.Service) (http.Handler, error) {
	warrantCfg, ok := cfg.(config.Forge4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Forge4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, tokenString, err := service.ParseAuthTokenFromRequest(r, []string{service.AuthTypeApiKey})
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		if !service.SecureCompareEqual(tokenString, warrantCfg.GetAuthentication().ApiKey) {
			service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key: Master API Key Required"))
			return
		}

		newContext := context.WithValue(r.Context(), service.AuthInfoKey, &service.AuthInfo{})
		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}

func MasterAndApiKeyAuthMiddleware(cfg config.Config, next http.Handler, svcs ...service.Service) (http.Handler, error) {
	var apiKeyService *ApiService
	for _, svc := range svcs {
		if svc.ID() == service.ApiService {
			apiKeyService = svc.(*ApiService)
		}
	}

	warrantCfg, ok := cfg.(config.Forge4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Forge4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, tokenString, err := service.ParseAuthTokenFromRequest(r, []string{service.AuthTypeApiKey})
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		newContext := context.WithValue(r.Context(), service.AuthInfoKey, &service.AuthInfo{})

		if service.SecureCompareEqual(tokenString, warrantCfg.GetAuthentication().ApiKey) {
			next.ServeHTTP(w, r.WithContext(newContext))
			return
		}

		dbKey, err := apiKeyService.GetByKey(context.Background(), tokenString)
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key"))
			return
		}

		if !service.SecureCompareEqual(tokenString, dbKey.ApiKey) {
			service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key"))
			return
		}

		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}
