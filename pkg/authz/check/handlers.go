package authz

import (
	"net/http"

	session "github.com/forge4flow/forge4flow-core/pkg/authn/session"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
	warrant "github.com/forge4flow/forge4flow-core/pkg/authz/warrant"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc CheckService) Routes() ([]service.Route, error) {
	return []service.Route{
		// Standard Authorization
		service.ForgeRoute{
			Pattern:                    "/v2/authorize",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, AuthorizeHandler),
			OverrideAuthMiddlewareFunc: session.ApiKeyAndSessionAuthMiddleware,
		},
	}, nil
}

func AuthorizeHandler(svc CheckService, w http.ResponseWriter, r *http.Request) error {
	authInfo := service.GetAuthInfoFromRequestContext(r.Context())
	if authInfo != nil && authInfo.UserId != "" {
		var sessionCheckManySpec SessionCheckManySpec
		err := service.ParseJSONBody(r.Body, &sessionCheckManySpec)
		if err != nil {
			return err
		}

		warrantSpecs := make([]CheckWarrantSpec, 0)
		for _, warrantSpec := range sessionCheckManySpec.Warrants {
			warrantSpecs = append(warrantSpecs, CheckWarrantSpec{
				ObjectType: warrantSpec.ObjectType,
				ObjectId:   warrantSpec.ObjectId,
				Relation:   warrantSpec.Relation,
				Subject: &warrant.SubjectSpec{
					ObjectType: objecttype.ObjectTypeUser,
					ObjectId:   authInfo.UserId,
				},
			})
		}

		checkManySpec := CheckManySpec{
			Op:       sessionCheckManySpec.Op,
			Warrants: warrantSpecs,
			Context:  sessionCheckManySpec.Context,
			Debug:    sessionCheckManySpec.Debug,
		}

		checkResult, err := svc.CheckMany(r.Context(), authInfo, &checkManySpec)
		if err != nil {
			return err
		}

		service.SendJSONResponse(w, checkResult)
		return nil
	}

	var checkManySpec CheckManySpec
	err := service.ParseJSONBody(r.Body, &checkManySpec)
	if err != nil {
		return err
	}

	checkResult, err := svc.CheckMany(r.Context(), authInfo, &checkManySpec)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, checkResult)
	return nil
}
