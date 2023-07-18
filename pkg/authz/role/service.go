package authz

import (
	"context"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
	permission "github.com/forge4flow/forge4flow-core/pkg/authz/permission"
	warrant "github.com/forge4flow/forge4flow-core/pkg/authz/warrant"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

const ResourceTypeRole = "role"

type RoleService struct {
	service.BaseService
	Repository    RoleRepository
	WarrantSvc    *warrant.WarrantService
	PermissionSvc *permission.PermissionService
	EventSvc      *event.EventService
	ObjectSvc     *object.ObjectService
}

func NewService(env service.Env, repository RoleRepository, warrantSvc *warrant.WarrantService, permissionSvc *permission.PermissionService, eventSvc *event.EventService, objectSvc *object.ObjectService) *RoleService {
	return &RoleService{
		BaseService:   service.NewBaseService(env),
		Repository:    repository,
		WarrantSvc:    warrantSvc,
		PermissionSvc: permissionSvc,
		EventSvc:      eventSvc,
		ObjectSvc:     objectSvc,
	}
}

func (svc RoleService) ID() string {
	return service.RoleService
}

func (svc RoleService) Create(ctx context.Context, roleSpec RoleSpec) (*RoleSpec, error) {
	var newRole Model
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		_, err := svc.Repository.GetByRoleId(txCtx, roleSpec.RoleId)
		if err == nil {
			return service.NewDuplicateRecordError("Role", roleSpec.RoleId, "A role with the given roleId already exists")
		}

		createdObject, err := svc.ObjectSvc.Create(txCtx, *roleSpec.ToObjectSpec())
		if err != nil {
			return err
		}

		newRoleId, err := svc.Repository.Create(txCtx, roleSpec.ToRole(createdObject.ID))
		if err != nil {
			return err
		}

		newRole, err = svc.Repository.GetById(txCtx, newRoleId)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypeRole, newRole.GetRoleId(), newRole.ToRoleSpec())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newRole.ToRoleSpec(), nil
}

func (svc RoleService) GetByRoleId(ctx context.Context, roleId string) (*RoleSpec, error) {
	role, err := svc.Repository.GetByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	return role.ToRoleSpec(), nil
}

func (svc RoleService) List(ctx context.Context, listParams service.ListParams) ([]RoleSpec, error) {
	roleSpecs := make([]RoleSpec, 0)

	roles, err := svc.Repository.List(ctx, listParams)
	if err != nil {
		return roleSpecs, nil
	}

	for _, role := range roles {
		roleSpecs = append(roleSpecs, *role.ToRoleSpec())
	}

	return roleSpecs, nil
}

func (svc RoleService) UpdateByRoleId(ctx context.Context, roleId string, roleSpec UpdateRoleSpec) (*RoleSpec, error) {
	currentRole, err := svc.Repository.GetByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	currentRole.SetName(roleSpec.Name)
	currentRole.SetDescription(roleSpec.Description)
	err = svc.Repository.UpdateByRoleId(ctx, roleId, currentRole)
	if err != nil {
		return nil, err
	}

	updatedRole, err := svc.Repository.GetByRoleId(ctx, roleId)
	if err != nil {
		return nil, err
	}

	updatedRoleSpec := updatedRole.ToRoleSpec()
	err = svc.EventSvc.TrackResourceUpdated(ctx, ResourceTypeRole, updatedRole.GetRoleId(), updatedRoleSpec)
	if err != nil {
		return nil, err
	}

	return updatedRoleSpec, nil
}

func (svc RoleService) DeleteByRoleId(ctx context.Context, roleId string) error {
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		err := svc.Repository.DeleteByRoleId(txCtx, roleId)
		if err != nil {
			return err
		}

		err = svc.ObjectSvc.DeleteByObjectTypeAndId(txCtx, objecttype.ObjectTypeRole, roleId)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceDeleted(txCtx, ResourceTypeRole, roleId, nil)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc RoleService) GetAllObjectsForRoleByType(ctx context.Context, roleId string, objectType string) ([]object.ObjectSpec, error) {
	matchingObjects := make([]object.ObjectSpec, 0)
	matchingWarrants, err := svc.WarrantSvc.GetAllWarrantsForSubjectIdByType(ctx, "role", roleId, objectType)
	if err != nil {
		return matchingObjects, err
	}
	// Convert warrants to ObjectSpec
	for _, matchingWarrant := range matchingWarrants {
		object, err := svc.ObjectSvc.GetByObjectTypeAndId(ctx, matchingWarrant.GetObjectType(), matchingWarrant.GetObjectId())
		if err != nil {
			return matchingObjects, err
		}

		objectSpec := object.ToObject().ToObjectSpec()

		matchingObjects = append(matchingObjects, *objectSpec)
	}

	return matchingObjects, nil
}
