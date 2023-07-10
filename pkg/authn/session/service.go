package authn

import (
	"context"

	nonce "github.com/auth4flow/auth4flow-core/pkg/authn/nonce"
	user "github.com/auth4flow/auth4flow-core/pkg/authz/user"
	"github.com/auth4flow/auth4flow-core/pkg/config"
	"github.com/auth4flow/auth4flow-core/pkg/event"
	"github.com/auth4flow/auth4flow-core/pkg/service"
)

type SessionService struct {
	service.BaseService
	Config     config.Auth4FlowConfig
	Repository SessionRepository
	NonceSvc   *nonce.NonceService
	UserSvc    *user.UserService
	EventSvc   *event.EventService
}

func NewService(env service.Env, cfg config.Auth4FlowConfig, repo SessionRepository, nonceSvc *nonce.NonceService, userSvc *user.UserService, eventSvc *event.EventService) *SessionService {
	return &SessionService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		Repository:  repo,
		NonceSvc:    nonceSvc,
		UserSvc:     userSvc,
		EventSvc:    eventSvc,
	}
}

func (svc SessionService) ID() string {
	return service.SessionService
}

func (svc SessionService) Create(ctx context.Context, sessionDetails SessionCreationSpec) (*SessionSpec, error) {
	var newSession Model
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		newNonceId, err := svc.Repository.Create(txCtx, sessionDetails.ToSession())
		if err != nil {
			return err
		}

		newSession, err = svc.Repository.GetById(txCtx, newNonceId)
		if err != nil {
			return err
		}

		//TODO: Update for proper logging of session creation
		// err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypeUser, newUser.GetUserId(), newUser.ToUserSpec())
		// if err != nil {
		// 	return err
		// }

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newSession.ToSessionSpec(), nil
}
