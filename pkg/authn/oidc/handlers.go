package authn

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

// TODO: NONE OF THIS SHOULD BE HARD CODED!! - TESTING ONLY
var GOOGLE = OauthProvider{
	ID:           "google",
	Name:         "Google",
	ImgURL:       "",
	UserIdClaim:  "sub",
	Scopes:       []string{"openid", "email"},
	ClientID:     "986980395964-50o8blm4pfftfea0k81ik4vk7b7rbcfj.apps.googleusercontent.com",
	ClientSecret: "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",
	CodeURL:      "https://accounts.google.com/o/oauth2/v2/auth",
	TokenURL:     "https://oauth2.googleapis.com/token",
}

func (svc OauthService) Routes() ([]service.Route, error) {
	return []service.Route{
		service.WarrantRoute{
			Pattern:                    "/v1/oidc/providers",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, GetProvidersHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oidc/configuration",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, GetProvidersHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oidc/provider",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, GetProvidersHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/oidc/provider",
			Method:                     "PUT",
			Handler:                    service.NewRouteHandler(svc, GetProvidersHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},
	}, nil
}

func GetProvidersHandler(svc OauthService, w http.ResponseWriter, r *http.Request) error {
	providers := []OauthProvider{
		GOOGLE,
	}

	service.SendJSONResponse(w, providers)
	return nil
}
