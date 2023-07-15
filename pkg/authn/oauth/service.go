package authn

import (
	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

type OauthService struct {
	service.BaseService
	Config   config.Forge4FlowConfig
	EventSvc *event.EventService
}

func NewService(env service.Env, cfg config.Forge4FlowConfig, eventSvc *event.EventService) *OauthService {
	return &OauthService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		EventSvc:    eventSvc,
	}
}

func (svc OauthService) ID() string {
	return service.NonceService
}
