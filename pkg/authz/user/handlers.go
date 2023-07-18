package authz

import (
	"net/http"

	permission "github.com/forge4flow/forge4flow-core/pkg/authz/permission"
	role "github.com/forge4flow/forge4flow-core/pkg/authz/role"
	tenant "github.com/forge4flow/forge4flow-core/pkg/authz/tenant"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/gorilla/mux"
)

func (svc UserService) Routes() ([]service.Route, error) {
	return []service.Route{
		// create
		service.WarrantRoute{
			Pattern: "/v1/users",
			Method:  "POST",
			Handler: service.NewRouteHandler(svc, CreateHandler),
		},

		// get
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}",
			Method:  "GET",
			Handler: service.NewRouteHandler(svc, GetHandler),
		},
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}/tenants",
			Method:  "GET",
			Handler: service.NewRouteHandler(svc, GetTenantsForUserHandler),
		},
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}/roles",
			Method:  "GET",
			Handler: service.NewRouteHandler(svc, GetRolesForUserHandler),
		},
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}/permissions",
			Method:  "GET",
			Handler: service.NewRouteHandler(svc, GetPermissionsForUserHandler),
		},
		service.WarrantRoute{
			Pattern: "/v1/users",
			Method:  "GET",
			Handler: service.ChainMiddleware(
				service.NewRouteHandler(svc, ListHandler),
				service.ListMiddleware[UserListParamParser],
			),
		},

		// delete
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}",
			Method:  "DELETE",
			Handler: service.NewRouteHandler(svc, DeleteHandler),
		},

		// update
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}",
			Method:  "POST",
			Handler: service.NewRouteHandler(svc, UpdateHandler),
		},
		service.WarrantRoute{
			Pattern: "/v1/users/{userId}",
			Method:  "PUT",
			Handler: service.NewRouteHandler(svc, UpdateHandler),
		},
	}, nil
}

func CreateHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	var userSpec UserSpec
	err := service.ParseJSONBody(r.Body, &userSpec)
	if err != nil {
		return err
	}

	createdUser, err := svc.Create(r.Context(), userSpec)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, createdUser)
	return nil
}

func GetHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	user, err := svc.GetByUserId(r.Context(), userId)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, user)
	return nil
}

func GetTenantsForUserHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	tenantObjects, err := svc.GetAllObjectsForUserByType(r.Context(), userId, "tenant")
	if err != nil {
		return err
	}

	tenants := make([]*tenant.TenantSpec, 0)
	for _, tenantObject := range tenantObjects {
		tenant, err := svc.TenantSvc.GetByTenantId(r.Context(), tenantObject.ObjectId)
		if err != nil {
			return err
		}

		tenants = append(tenants, tenant)
	}

	service.SendJSONResponse(w, tenants)

	return nil
}

func GetRolesForUserHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	roleObjects, err := svc.GetAllObjectsForUserByType(r.Context(), userId, "role")
	if err != nil {
		return err
	}

	roles := make([]*role.RoleSpec, 0)
	for _, roleObject := range roleObjects {
		role, err := svc.RoleSvc.GetByRoleId(r.Context(), roleObject.ObjectId)
		if err != nil {
			return err
		}

		roles = append(roles, role)
	}

	service.SendJSONResponse(w, roles)

	return nil
}

func GetPermissionsForUserHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	permissionObjects, err := svc.GetAllObjectsForUserByType(r.Context(), userId, "permission")
	if err != nil {
		return err
	}

	permissions := make([]*permission.PermissionSpec, 0)
	for _, permissionObject := range permissionObjects {
		permission, err := svc.PermissionSvc.GetByPermissionId(r.Context(), permissionObject.ObjectId)
		if err != nil {
			return err
		}

		permissions = append(permissions, permission)
	}

	service.SendJSONResponse(w, permissions)

	return nil
}

func ListHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	listParams := service.GetListParamsFromContext(r.Context())
	users, err := svc.List(r.Context(), listParams)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, users)
	return nil
}

func UpdateHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	var updateUser UpdateUserSpec
	err := service.ParseJSONBody(r.Body, &updateUser)
	if err != nil {
		return err
	}

	userId := mux.Vars(r)["userId"]
	updatedUser, err := svc.UpdateByUserId(r.Context(), userId, updateUser)
	if err != nil {
		return err
	}

	service.SendJSONResponse(w, updatedUser)
	return nil
}

func DeleteHandler(svc UserService, w http.ResponseWriter, r *http.Request) error {
	userId := mux.Vars(r)["userId"]
	err := svc.DeleteByUserId(r.Context(), userId)
	if err != nil {
		return err
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	return nil
}
