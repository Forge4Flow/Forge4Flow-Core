package authz

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

// GetRoutes registers all route handlers for this module
func (svc WarrantService) Routes() ([]service.Route, error) {
	return []service.Route{
		// create
		service.ForgeRoute{
			Pattern: "/v1/warrants",
			Method:  "POST",
			Handler: service.ChainMiddleware(
				service.NewRouteHandler(svc, CreateHandler),
			),
		},

		// list
		service.ForgeRoute{
			Pattern: "/v1/warrants",
			Method:  "GET",
			Handler: service.ChainMiddleware(
				service.NewRouteHandler(svc, ListHandler),
				service.ListMiddleware[WarrantListParamParser],
			),
		},

		// delete
		service.ForgeRoute{
			Pattern: "/v1/warrants",
			Method:  "DELETE",
			Handler: service.ChainMiddleware(
				service.NewRouteHandler(svc, DeleteHandler),
			),
		},
	}, nil
}

func CreateHandler(svc WarrantService, w http.ResponseWriter, r *http.Request) error {
	var warrantSpec WarrantSpec
	err := service.ParseJSONBody(r.Body, &warrantSpec)
	if err != nil {
		return err
	}

	// TODO: move into a custom golang-validate function
	if warrantSpec.Policy != "" {
		err := warrantSpec.Policy.Validate()
		if err != nil {
			return service.NewInvalidParameterError("policy", err.Error())
		}
	}

	createdWarrant, err := svc.Create(r.Context(), warrantSpec)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, createdWarrant)
	return nil
}

func ListHandler(svc WarrantService, w http.ResponseWriter, r *http.Request) error {
	listParams := service.GetListParamsFromContext(r.Context())
	queryParams := r.URL.Query()
	filters := FilterOptions{
		ObjectType: queryParams.Get("objectType"),
		ObjectId:   queryParams.Get("objectId"),
		Relation:   queryParams.Get("relation"),
		Subject: &SubjectSpec{
			ObjectType: queryParams.Get("subjectType"),
			ObjectId:   queryParams.Get("subjectId"),
		},
		Policy: Policy(queryParams.Get("policy")),
	}
	subjectRelation := queryParams.Get("subjectRelation")
	if subjectRelation != "" {
		filters.Subject.Relation = subjectRelation
	}

	warrants, err := svc.List(r.Context(), &filters, listParams)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, warrants)
	return nil
}

func DeleteHandler(svc WarrantService, w http.ResponseWriter, r *http.Request) error {
	var warrantSpec WarrantSpec
	err := service.ParseJSONBody(r.Body, &warrantSpec)
	if err != nil {
		return err
	}

	err = svc.Delete(r.Context(), warrantSpec)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
