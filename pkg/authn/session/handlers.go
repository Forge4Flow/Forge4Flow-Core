package authn

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	user "github.com/forge4flow/forge4flow-core/pkg/authz/user"
	"github.com/forge4flow/forge4flow-core/pkg/flow"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc SessionService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Create Session
		service.WarrantRoute{
			Pattern:                    "/v1/session",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, CreateSessionHandler),
			OverrideAuthMiddlewareFunc: flow.ApiKeyAndAccountProofAuthMidleware,
		},

		// Verify Session
		service.WarrantRoute{
			Pattern:                    "/v1/session/verify",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, VerifySessionHandler),
			OverrideAuthMiddlewareFunc: SessionAuthMiddleware,
		},
		service.WarrantRoute{
			Pattern:                    "/v1/session/verify",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, VerifySessionHandler),
			OverrideAuthMiddlewareFunc: ApiKeyAndSessionAuthMiddleware,
		},
	}, nil
}

func CreateSessionHandler(svc SessionService, w http.ResponseWriter, r *http.Request) error {
	authInfo := service.GetAuthInfoFromRequestContext(r.Context())
	if authInfo != nil && authInfo.UserId != "" {
		// Verify that the user exists
		_, err := svc.UserSvc.GetByUserId(r.Context(), authInfo.UserId)
		if err != nil {
			if _, ok := err.(*service.RecordNotFoundError); ok {
				// If doesn't exist, and auto register enabled, then create user
				if svc.Config.GetAuthentication().AutoRegister {
					newUser := user.UserSpec{
						UserId: authInfo.UserId,
					}
					_, err := svc.UserSvc.Create(r.Context(), newUser)
					if err != nil {
						service.SendErrorResponse(w, service.NewInternalError("unable to register new user"))
						return err
					}
				} else {
					return err
				}
			} else {
				service.SendErrorResponse(w, service.NewInternalError("error locating user record"))
				return err
			}
		}

		// Create Session ID/Token
		sessionToken, err := generateSessionToken(svc.Config.Authentication.SessionTokenLength)
		if err != nil {
			service.SendErrorResponse(w, service.NewInternalError("failed to create session token"))
			return err
		}

		// Create Session Details
		sessionDetails := SessionCreationSpec{
			SessionId:   sessionToken,
			UserId:      authInfo.UserId,
			IdleTimeout: svc.Config.GetAuthentication().SessionIdleTimeout,
			ExpTime:     time.Now().Add(time.Duration(svc.Config.GetAuthentication().SessionExpTimeout)).UTC(),
			ClientIp:    service.GetClientIpAddress(r),
			UserAgent:   r.UserAgent(),
		}

		session, err := svc.Create(r.Context(), sessionDetails)
		if err != nil {
			service.SendErrorResponse(w, service.NewInternalError("failed to create session"))
			return err
		}

		service.SendJSONResponse(w, session)
		return nil
	}

	// Invalid authInfo in Context - Return UnAuthorizedError
	service.SendErrorResponse(w, service.NewUnauthorizedError("Invalid Auth Info"))
	return nil
}

func generateSessionToken(tokenLength int64) (string, error) {
	tokenBytes := make([]byte, tokenLength)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes as a base64 string to make it readable and usable as a session token
	sessionToken := base64.URLEncoding.EncodeToString(tokenBytes)
	return sessionToken, nil
}

func VerifySessionHandler(svc SessionService, w http.ResponseWriter, r *http.Request) error {
	authInfo := service.GetAuthInfoFromRequestContext(r.Context())
	if authInfo != nil && authInfo.UserId != "" {
		validSession := &SessionVerificationSpec{
			UserId: authInfo.UserId,
			Result: "Valid",
		}
		service.SendJSONResponse(w, validSession)
		return nil
	}

	service.SendErrorResponse(w, service.NewInternalError("unable to validate session"))
	return nil
}
