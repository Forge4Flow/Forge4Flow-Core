package authn

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc ApiService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Create API Key
		service.ForgeRoute{
			Pattern:                    "/v1/api",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, CreateApiKey),
			OverrideAuthMiddlewareFunc: MasterApiKeyAuthMiddleware,
		},

		// Delete API Key
		service.ForgeRoute{
			Pattern:                    "/v1/api",
			Method:                     "DELETE",
			Handler:                    service.NewRouteHandler(svc, DeleteApiKey),
			OverrideAuthMiddlewareFunc: MasterApiKeyAuthMiddleware,
		},
	}, nil
}

func CreateApiKey(svc ApiService, w http.ResponseWriter, r *http.Request) error {
	var request ApiSpec
	err := service.ParseJSONBody(r.Body, &request)
	if err != nil {
		return err
	}

	newApiKey, err := svc.Create(r.Context(), request.DisplayName)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, newApiKey)
	return nil
}

func DeleteApiKey(svc ApiService, w http.ResponseWriter, r *http.Request) error {
	var request ApiSpec
	err := service.ParseJSONBody(r.Body, &request)
	if err != nil {
		return err
	}

	err = svc.DeleteByKey(r.Context(), request.ApiKey)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
