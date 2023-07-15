package authn

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc OauthService) Routes() ([]service.Route, error) {
	return []service.Route{
		service.WarrantRoute{
			Pattern:                    "/v1/oauth",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, OauthRedirectHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oauth/callback",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, OauthRedirectHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oauth/providers",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, GetProvidersHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oauth/configuration",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, OauthRedirectHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oauth/provider",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, OauthRedirectHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oauth/provider",
			Method:                     "PUT",
			Handler:                    service.NewRouteHandler(svc, OauthRedirectHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},
	}, nil
}

func GetProvidersHandler(svc OauthService, w http.ResponseWriter, r *http.Request) error {
	providers := []OauthProvider{
		{
			ID:     "google",
			Name:   "Google",
			ImgURL: "",
		},
	}

	service.SendJSONResponse(w, providers)
	return nil
}

func OauthRedirectHandler(svc OauthService, w http.ResponseWriter, r *http.Request) error {
	// Get Provider from query parameters
	queryParams := r.URL.Query()
	if !queryParams.Has("provider") {
		service.SendErrorResponse(w, service.NewMissingRequiredParameterError("provider"))
		return nil
	}
	provider := queryParams.Get("provider")

	// Verify profider is configured & Get details from db

	// build redirect URL

	// redirect to provider endpoint

	service.SendErrorResponse(w, service.NewInternalError("not conifgured"))
	return nil
}
