package authn

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/auth4flow/auth4flow-core/pkg/config"
	"github.com/auth4flow/auth4flow-core/pkg/event"
	"github.com/auth4flow/auth4flow-core/pkg/service"
)

type NonceService struct {
	service.BaseService
	Config     config.Auth4FlowConfig
	Repository NonceRepository
	EventSvc   *event.EventService
}

func NewService(env service.Env, cfg config.Auth4FlowConfig, nonceRepo NonceRepository, eventSvc *event.EventService) *NonceService {
	return &NonceService{
		BaseService: service.NewBaseService(env),
		Config:      cfg,
		Repository:  nonceRepo,
		EventSvc:    eventSvc,
	}
}

func (svc NonceService) ID() string {
	return service.NonceService
}

func (svc NonceService) Create(ctx context.Context) (*NonceSpec, error) {
	//Generate nonce and nonceSpec
	nonce, err := generateNonce()
	if err != nil {
		return nil, err
	}

	nonceSpec := &NonceSpec{
		Nonce: nonce,
	}

	var newNonce Model
	err = svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		newNonceId, err := svc.Repository.Create(txCtx, nonceSpec.ToNonce())
		if err != nil {
			return err
		}

		newNonce, err = svc.Repository.GetById(txCtx, newNonceId)
		if err != nil {
			return err
		}

		//TODO: Update for proper logging of nonce creation
		// err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypeUser, newUser.GetUserId(), newUser.ToUserSpec())
		// if err != nil {
		// 	return err
		// }

		return nil
	})
	if err != nil {
		return nil, err
	}

	newNonceSpec := newNonce.ToNonceSpec()

	newNonceSpec.AppIdentifier = svc.Config.GetAppIdentifier()

	return newNonceSpec, nil
}

func (svc NonceService) IsValid(ctx context.Context, nonce string) (bool, error) {
	validNonce := false
	nonceModel, err := svc.Repository.GetByNonce(ctx, nonce)
	if err != nil {
		return validNonce, err
	}

	if !nonceModel.IsExpired() {
		validNonce = true
	}

	err = svc.Repository.DeleteByNonce(ctx, nonce)

	return validNonce, err
}

func generateNonce() (string, error) {
	nonceBytes := make([]byte, 32)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", fmt.Errorf("could not generate nonce")
	}

	return hex.EncodeToString(nonceBytes), nil
}
