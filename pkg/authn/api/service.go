package authn

import (
	"context"
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/forge4flow/forge4flow-core/utils"
)

type ApiService struct {
	service.BaseService
	Config     config.Forge4FlowConfig
	Repository ApiKeyRepository
	ObjectSvc  *object.ObjectService
	EventSvc   *event.EventService
}

const ResourceTypeApiKey = "apiKey"

func NewService(env service.Env, cfg config.Forge4FlowConfig, repo ApiKeyRepository, objectSvc *object.ObjectService, eventSvc *event.EventService) *ApiService {
	return &ApiService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		Repository:  repo,
		ObjectSvc:   objectSvc,
		EventSvc:    eventSvc,
	}
}

func (svc ApiService) ID() string {
	return service.ApiService
}

func (svc ApiService) Create(ctx context.Context, displayName string) (*ApiSpec, error) {
	//Generate apiKey and apiSpec
	key, err := utils.GenerateRandomKey()
	if err != nil {
		return nil, err
	}

	apiKey := &key

	apiSpec := &ApiSpec{
		DisplayName: displayName,
		Key:         apiKey,
		ExpDate:     time.Now().Add(time.Hour * 24 * 90),
	}

	var newApiKey Model
	err = svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		createdObject, err := svc.ObjectSvc.Create(txCtx, *apiSpec.ToObjectSpec())
		if err != nil {
			switch err.(type) {
			case *service.DuplicateRecordError:
				return service.NewDuplicateRecordError("API Key", apiSpec.Key, "An api key with the given key already exists")
			default:
				return err
			}
		}

		_, err = svc.Repository.GetByKey(txCtx, *apiSpec.Key)
		if err == nil {
			return service.NewDuplicateRecordError("API Key", apiSpec.Key, "An api key with the given key already exists")
		}

		newApiKeyId, err := svc.Repository.Create(txCtx, apiSpec.ToApiKey(createdObject.ID))
		if err != nil {
			return err
		}

		newApiKey, err = svc.Repository.GetById(txCtx, newApiKeyId)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypeApiKey, *newApiKey.GetKey(), newApiKey.ToApiSpec())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	newNonceSpec := newApiKey.ToApiSpec()

	return newNonceSpec, nil
}

func (svc ApiService) GetByKey(ctx context.Context, apiKey string) (*ApiSpec, error) {
	key, err := svc.Repository.GetByKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	return key.ToApiSpec(), nil
}

// TODO: List function
// func (svc ApiService) List(ctx context.Context, listParams service.ListParams) ([]ApiSpec, error) {
// 	ApiSpecs := make([]ApiSpec, 0)

// 	users, err := svc.Repository.List(ctx, listParams)
// 	if err != nil {
// 		return ApiSpecs, err
// 	}

// 	for _, user := range users {
// 		ApiSpecs = append(ApiSpecs, *user.ToApiSpec())
// 	}

// 	return ApiSpecs, nil
// }

func (svc ApiService) DeleteByKey(ctx context.Context, apiKey string) error {
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		err := svc.Repository.DeleteByKey(txCtx, apiKey)
		if err != nil {
			return err
		}

		err = svc.ObjectSvc.DeleteByObjectTypeAndId(txCtx, objecttype.ObjectTypeUser, apiKey)
		if err != nil {
			return err
		}

		err = svc.EventSvc.TrackResourceDeleted(txCtx, ResourceTypeApiKey, apiKey, nil)
		if err != nil {
			return err
		}

		return nil
	})

	return err
}
