package authn

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/auth4flow/auth4flow-core/pkg/event"
	"github.com/auth4flow/auth4flow-core/pkg/service"
)

type NonceService struct {
	service.BaseService
	Repository NonceRepository
	EventSvc   event.EventService
}

func NewService(env service.Env, nonceRepo NonceRepository, eventSvc event.EventService) NonceService {
	return NonceService{
		BaseService: service.NewBaseService(env),
		Repository:  nonceRepo,
		EventSvc:    eventSvc,
	}
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

	return newNonce.ToNonceSpec(), nil
}

func generateNonce() (string, error) {
	nonceBytes := make([]byte, 32)
	_, err := rand.Read(nonceBytes)
	if err != nil {
		return "", fmt.Errorf("could not generate nonce")
	}

	return base64.URLEncoding.EncodeToString(nonceBytes), nil
}
