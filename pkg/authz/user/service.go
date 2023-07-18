package authz

import (
	"context"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
	permission "github.com/forge4flow/forge4flow-core/pkg/authz/permission"
	role "github.com/forge4flow/forge4flow-core/pkg/authz/role"
	tenant "github.com/forge4flow/forge4flow-core/pkg/authz/tenant"
	warrant "github.com/forge4flow/forge4flow-core/pkg/authz/warrant"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

const ResourceTypeUser = "user"

type UserService struct {
	service.BaseService
	Repository    UserRepository
	PermissionSvc *permission.PermissionService
	RoleSvc       *role.RoleService
	TenantSvc     *tenant.TenantService
	WarrantSvc    *warrant.WarrantService
	EventSvc      *event.EventService
	ObjectSvc     *object.ObjectService
}

func NewService(env service.Env, repository UserRepository, permissionSvc *permission.PermissionService, roleSvc *role.RoleService, tenantSvc *tenant.TenantService, warrantSvc *warrant.WarrantService, eventSvc *event.EventService, objectSvc *object.ObjectService) *UserService {
	return &UserService{
		BaseService:   service.NewBaseService(env),
		Repository:    repository,
		PermissionSvc: permissionSvc,
		RoleSvc:       roleSvc,
		TenantSvc:     tenantSvc,
		WarrantSvc:    warrantSvc,
		EventSvc:      eventSvc,
		ObjectSvc:     objectSvc,
	}
}

func (svc UserService) ID() string {
	return service.UserService
}

func (svc UserService) Create(ctx context.Context, userSpec UserSpec) (*UserSpec, error) {
	if userSpec.UserId == "" {
		// generate an id for the user if one isn't provided
		generatedUUID, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.New("unable to generate random UUID for user")
		}
		userSpec.UserId = generatedUUID.String()
	}

	var newUser Model
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		createdObject, err := svc.ObjectSvc.Create(txCtx, *userSpec.ToObjectSpec())
		if err != nil {
			switch err.(type) {
			case *service.DuplicateRecordError:
				return service.NewDuplicateRecordError("User", userSpec.UserId, "A user with the given userId already exists")
			default:
				return err
			}
		}

		_, err = svc.Repository.GetByUserId(txCtx, userSpec.UserId)
		if err == nil {
			return service.NewDuplicateRecordError("User", userSpec.UserId, "A user with the given userId already exists")
		}

		newUserId, err := svc.Repository.Create(txCtx, userSpec.ToUser(createdObject.ID))
		if err != nil {
			return err
		}

		newUser, err = svc.Repository.GetById(txCtx, newUserId)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypeUser, newUser.GetUserId(), newUser.ToUserSpec())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newUser.ToUserSpec(), nil
}

func (svc UserService) GetByUserId(ctx context.Context, userId string) (*UserSpec, error) {
	user, err := svc.Repository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	return user.ToUserSpec(), nil
}

func (svc UserService) List(ctx context.Context, listParams service.ListParams) ([]UserSpec, error) {
	userSpecs := make([]UserSpec, 0)

	users, err := svc.Repository.List(ctx, listParams)
	if err != nil {
		return userSpecs, err
	}

	for _, user := range users {
		userSpecs = append(userSpecs, *user.ToUserSpec())
	}

	return userSpecs, nil
}

func (svc UserService) UpdateByUserId(ctx context.Context, userId string, userSpec UpdateUserSpec) (*UserSpec, error) {
	currentUser, err := svc.Repository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	currentUser.SetEmail(userSpec.Email)
	err = svc.Repository.UpdateByUserId(ctx, userId, currentUser)
	if err != nil {
		return nil, err
	}

	updatedUser, err := svc.Repository.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	updatedUserSpec := updatedUser.ToUserSpec()
	err = svc.EventSvc.TrackResourceUpdated(ctx, ResourceTypeUser, updatedUser.GetUserId(), updatedUserSpec)
	if err != nil {
		return nil, err
	}

	return updatedUserSpec, nil
}

func (svc UserService) DeleteByUserId(ctx context.Context, userId string) error {
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		err := svc.Repository.DeleteByUserId(txCtx, userId)
		if err != nil {
			return err
		}

		err = svc.ObjectSvc.DeleteByObjectTypeAndId(txCtx, objecttype.ObjectTypeUser, userId)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceDeleted(txCtx, ResourceTypeUser, userId, nil)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func (svc UserService) GetAllObjectsForUserByType(ctx context.Context, userId string, objectType string) ([]object.ObjectSpec, error) {
	matchingObjects := make([]object.ObjectSpec, 0)
	matchingWarrants, err := svc.WarrantSvc.GetAllWarrantsForSubjectIdByType(ctx, "user", userId, objectType)
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
