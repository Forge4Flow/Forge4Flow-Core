package authn

import (
	"context"
	"fmt"
	"net/http"

	"github.com/auth4flow/auth4flow-core/pkg/config"
	"github.com/auth4flow/auth4flow-core/pkg/service"
	"github.com/pkg/errors"
)

func SessionAuthMiddleware(cfg config.Config, next http.Handler, svcs ...service.Service) (http.Handler, error) {
	var sessionService *SessionService
	for _, svc := range svcs {
		if svc.ID() == service.SessionService {
			sessionService = svc.(*SessionService)
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger := hlog.FromRequest(r)
		_, tokenString, err := service.ParseAuthTokenFromRequest(r, []string{service.AuthTypeApiKey, service.AuthTypeBearer})
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		// Get Session
		session, err := sessionService.Repository.GetBySessionId(r.Context(), tokenString)
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid SessionId"))
			return
		}

		// Check If Session Has Expired
		if session.IsExpired() {
			sessionService.Repository.DeleteById(r.Context(), session.GetID())
			service.SendErrorResponse(w, service.NewTokenExpiredError())
			return
		}

		// Verify User Agent Matches
		if !service.SecureCompareEqual(r.UserAgent(), session.GetUserAgent()) {
			sessionService.Repository.DeleteById(r.Context(), session.GetID())
			service.SendErrorResponse(w, service.NewUnauthorizedError("User Agent Does Not Match"))
			return
		}

		// Update Session Activity
		err = sessionService.Repository.UpdateSessionActivity(r.Context(), session.GetID())
		if err != nil {
			sessionService.Repository.DeleteById(r.Context(), session.GetID())
			service.SendErrorResponse(w, service.NewUnauthorizedError("Error Updating Session Activity; Session Destroyed"))
			return
		}

		// Valid Session
		authInfo := &service.AuthInfo{
			UserId: session.GetUserId(),
		}

		newContext := context.WithValue(r.Context(), service.AuthInfoKey, *authInfo)
		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}

func ApiKeyAndSessionAuthMiddleware(cfg config.Config, next http.Handler, svcs ...service.Service) (http.Handler, error) {
	var sessionService *SessionService
	for _, svc := range svcs {
		if svc.ID() == service.SessionService {
			sessionService = svc.(*SessionService)
		}
	}

	auth4FlowConfig, ok := cfg.(config.Auth4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Auth4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger := hlog.FromRequest(r)
		tokenType, tokenString, err := service.ParseAuthTokenFromRequest(r, []string{service.AuthTypeApiKey, service.AuthTypeBearer})
		if err != nil {
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		var authInfo *service.AuthInfo
		switch tokenType {
		case service.AuthTypeApiKey:
			if !service.SecureCompareEqual(tokenString, auth4FlowConfig.GetAuthentication().ApiKey) {
				service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key"))
				return
			}

			authInfo = &service.AuthInfo{}
		case service.AuthTypeBearer:
			// Get Session
			session, err := sessionService.Repository.GetBySessionId(r.Context(), tokenString)
			if err != nil {
				service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid SessionId"))
				return
			}

			// Check If Session Has Expired
			if session.IsExpired() {
				sessionService.Repository.DeleteById(r.Context(), session.GetID())
				service.SendErrorResponse(w, service.NewUnauthorizedError("Session Expired"))
				return
			}

			// Verify User Agent Matches
			if !service.SecureCompareEqual(r.UserAgent(), session.GetUserAgent()) {
				sessionService.Repository.DeleteById(r.Context(), session.GetID())
				service.SendErrorResponse(w, service.NewUnauthorizedError("User Agent Does Not Match"))
				return
			}

			// Update Session Activity
			err = sessionService.Repository.UpdateSessionActivity(r.Context(), session.GetID())
			if err != nil {
				sessionService.Repository.DeleteById(r.Context(), session.GetID())
				service.SendErrorResponse(w, service.NewUnauthorizedError("Error Updating Session Activity; Session Destroyed"))
				return
			}

			// Valid Session
			authInfo = &service.AuthInfo{
				UserId: session.GetUserId(),
			}
		}

		newContext := context.WithValue(r.Context(), service.AuthInfoKey, *authInfo)
		next.ServeHTTP(w, r.WithContext(newContext))
	}), nil
}
