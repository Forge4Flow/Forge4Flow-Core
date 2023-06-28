package authn

import (
	"net/http"

	"github.com/auth4flow/auth4flow-core/pkg/service"
)

func (svc SessionService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Create Session
		service.WarrantRoute{
			Pattern:                    "/v1/session",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, CreateSessionHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAndSessionAuthMiddleware,
		},
	}, nil
}

func CreateSessionHandler(svc SessionService, w http.ResponseWriter, r *http.Request) error {
	return nil
}
