package authn

import (
	"context"
	"time"

	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/forge4flow/forge4flow-core/utils"
)

type ApiService struct {
	service.BaseService
	Config     config.Forge4FlowConfig
	Repository ApiKeyRepository
	EventSvc   *event.EventService
}

const ResourceTypeApiKey = "apiKey"

func NewService(env service.Env, cfg config.Forge4FlowConfig, repo ApiKeyRepository, eventSvc *event.EventService) *ApiService {
	return &ApiService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		Repository:  repo,
		EventSvc:    eventSvc,
	}
}

func (svc ApiService) ID() string {
	return service.ApiService
}

func (svc ApiService) Create(ctx context.Context, displayName string, apiKey *string) (*ApiSpec, error) {
	//Generate apiKey and apiSpec
	if apiKey == nil {
		key, err := utils.GenerateRandomKey()
		if err != nil {
			return nil, err
		}

		apiKey = &key
	}

	apiSpec := &ApiSpec{
		DisplayName: displayName,
		Key:         *apiKey,
		ExpDate:     time.Now().Add(time.Hour * 24 * 90),
	}

	var newApiKey Model
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		newApiKeyId, err := svc.Repository.Create(txCtx, apiSpec.ToApiKey())
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
