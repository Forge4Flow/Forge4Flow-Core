package authn

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc NonceService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Get Nonce
		service.WarrantRoute{
			Pattern:                    "/v1/nonce",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, GenerateNonceHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},
	}, nil
}

func GenerateNonceHandler(svc NonceService, w http.ResponseWriter, r *http.Request) error {
	newNonce, err := svc.Create(r.Context())
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, newNonce)
	return nil
}
