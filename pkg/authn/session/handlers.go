package authn

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"github.com/auth4flow/auth4flow-core/pkg/service"
)

func (svc SessionService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Create Session
		service.WarrantRoute{
			Pattern:                    "/v1/session",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, CreateSessionHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAndAccountProofAuthMidleware,
		},
	}, nil
}

func CreateSessionHandler(svc SessionService, w http.ResponseWriter, r *http.Request) error {
	authInfo := service.GetAuthInfoFromRequestContext(r.Context())
	if authInfo != nil && authInfo.UserId != "" {
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
			IdleTimeout: time.Duration(svc.Config.Authentication.SessionIdleTimeout),
			ExpTime:     time.Now().Add(time.Duration(svc.Config.Authentication.SessionExpTimeout)),
			ClientIp:    service.GetClientIpAddress(r),
			UserAgent:   r.UserAgent(),
		}

		session, err := svc.Create(r.Context(), sessionDetails)
		if err != nil {
			service.SendErrorResponse(w, service.NewInternalError("failed to create session"))
			return err
		}

		service.SendJSONResponse(w, session)
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
