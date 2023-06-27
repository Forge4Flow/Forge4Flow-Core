package authn

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/auth4flow/auth4flow-core/pkg/service"
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
	nonceBytes := make([]byte, 64)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return fmt.Errorf("could not generate nonce")
	}

	nonceStruct := &NonceSpec{
		Nonce: base64.URLEncoding.EncodeToString(nonceBytes),
	}

	service.SendJSONResponse(w, nonceStruct)
	return nil
}
