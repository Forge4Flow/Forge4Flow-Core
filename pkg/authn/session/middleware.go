package authn

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/service"
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

		// TODO: FIX SESSION EXPERATION CHECKS
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

	forge4FlowConfig, ok := cfg.(config.Forge4FlowConfig)
	if !ok {
		return nil, errors.New("cfg parameter on DefaultAuthMiddleware must be a Forge4FlowConfig")
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logger := hlog.FromRequest(r)
		tokenType, tokenString, err := service.ParseAuthTokenFromRequest(r, []string{service.AuthTypeApiKey, service.AuthTypeBearer})
		if err != nil {
			fmt.Println("failed at checking token type")
			service.SendErrorResponse(w, service.NewUnauthorizedError(fmt.Sprintf("Invalid authorization header: %s", err.Error())))
			return
		}

		var authInfo *service.AuthInfo
		switch tokenType {
		case service.AuthTypeApiKey:
			if !service.SecureCompareEqual(tokenString, forge4FlowConfig.GetAuthentication().ApiKey) {
				fmt.Println("failed at checking api key")
				service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid API key"))
				return
			}

			authInfo = &service.AuthInfo{}
			if r.URL.Path == "/v1/session/verify" {
				var sessionSpec SessionVerificationSpec
				err := service.ParseJSONBody(r.Body, &sessionSpec)
				if err != nil {
					service.SendErrorResponse(w, err)
					return
				}

				// Get Session
				sessionId, err := url.QueryUnescape(sessionSpec.SessionId)
				if err != nil {
					service.SendErrorResponse(w, service.NewInternalError("Invalid session encoding"))
					return
				}

				session, err := sessionService.Repository.GetBySessionId(r.Context(), sessionId)
				if err != nil {
					fmt.Println(err)
					service.SendErrorResponse(w, service.NewInvalidRequestError("invalid session request"))
					return
				}

				authInfo.UserId = session.GetUserId()
			}
		case service.AuthTypeBearer:
			// Get Session
			session, err := sessionService.Repository.GetBySessionId(r.Context(), tokenString)
			if err != nil {
				fmt.Println("failed to get session")
				service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid SessionId"))
				return
			}

			// TODO: FIX SESSION EXPERATION CHECKS
			// Check If Session Has Expired
			if session.IsExpired() {
				sessionService.Repository.DeleteById(r.Context(), session.GetID())
				service.SendErrorResponse(w, service.NewTokenExpiredError())
				return
			}

			// Verify User Agent Matches
			if !service.SecureCompareEqual(r.UserAgent(), session.GetUserAgent()) {
				fmt.Println("failed to verify user agent")
				sessionService.Repository.DeleteById(r.Context(), session.GetID())
				service.SendErrorResponse(w, service.NewUnauthorizedError("User Agent Does Not Match"))
				return
			}

			// Update Session Activity
			err = sessionService.Repository.UpdateSessionActivity(r.Context(), session.GetID())
			if err != nil {
				fmt.Println("failed to update session activity")
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
