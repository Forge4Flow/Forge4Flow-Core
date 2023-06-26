package authn

import (
	"net/http"

	"github.com/auth4flow/auth4flow-core/pkg/service"
)

func (svc AuthnService) Routes() ([]service.Route, error) {
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

func GenerateNonceHandler(svc AuthnService, w http.ResponseWriter, r *http.Request) error {
	// TODO: Generate Nonce
	nonce := "12345678909876543212345678987654321"

	nonceStruct := &NonceSpec{
		Nonce: nonce,
	}

	service.SendJSONResponse(w, nonceStruct)
	return nil
}
