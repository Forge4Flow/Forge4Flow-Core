package authn

import (
	"github.com/auth4flow/auth4flow-core/pkg/event"
	"github.com/auth4flow/auth4flow-core/pkg/service"
)

type NonceService struct {
	service.BaseService
	NonceRepo NonceRepository
	EventSvc  event.EventService
}

func NewService(env service.Env, nonceRepo NonceRepository, eventSvc event.EventService) NonceService {
	return NonceService{
		BaseService: service.NewBaseService(env),
		NonceRepo:   nonceRepo,
		EventSvc:    eventSvc,
	}
}
