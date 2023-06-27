package authn

import (
	"github.com/auth4flow/auth4flow-core/pkg/event"
	"github.com/auth4flow/auth4flow-core/pkg/service"
)

type SessionService struct {
	service.BaseService
	EventSvc event.EventService
}

func NewService(env service.Env, eventSvc event.EventService) SessionService {
	return SessionService{
		BaseService: service.NewBaseService(env),
		EventSvc:    eventSvc,
	}
}
