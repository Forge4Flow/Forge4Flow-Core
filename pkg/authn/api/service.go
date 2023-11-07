package authn

import (
	"github.com/forge4flow/forge4flow-core/pkg/config"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

type ApiService struct {
	service.BaseService
	Config   config.Forge4FlowConfig
	EventSvc *event.EventService
}

func NewService(env service.Env, cfg config.Forge4FlowConfig, eventSvc *event.EventService) *ApiService {
	return &ApiService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		EventSvc:    eventSvc,
	}
}

func (svc ApiService) ID() string {
	return service.ApiService
}
