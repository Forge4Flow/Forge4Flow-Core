package authn

import (
	"time"
)

type Model interface {
	GetID() int64
	GetNonce() *string
	GetExpDate() time.Time
	IsExpired() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
	ToNonceSpec() *NonceSpec
}

type Nonce struct {
	ID        int64      `mysql:"id" postgres:"id" sqlite:"id"`
	Nonce     string     `mysql:"nonce" postgres:"nonce" sqlite:"nonce"`
	ExpDate   time.Time  `mysql:"expDate" postgres:"exp_date" sqlite:"expDate"`
	CreatedAt time.Time  `mysql:"createdAt" postgres:"created_at" sqlite:"createdAt"`
	UpdatedAt time.Time  `mysql:"updatedAt" postgres:"updated_at" sqlite:"updatedAt"`
	DeletedAt *time.Time `mysql:"deletedAt" postgres:"deleted_at" sqlite:"deletedAt"`
}

func (nonce Nonce) GetID() int64 {
	return nonce.ID
}

func (nonce Nonce) GetNonce() *string {
	if nonce.IsExpired() {
		return nil
	} else {
		return &nonce.Nonce
	}
}

func (nonce Nonce) GetExpDate() time.Time {
	return nonce.ExpDate
}

func (nonce Nonce) IsExpired() bool {
	return nonce.ExpDate.Before(time.Now())
}

func (nonce Nonce) GetCreatedAt() time.Time {
	return nonce.CreatedAt
}

func (nonce Nonce) GetUpdatedAt() time.Time {
	return nonce.UpdatedAt
}

func (nonce Nonce) GetDeletedAt() *time.Time {
	return nonce.DeletedAt
}

func (nonce Nonce) ToNonceSpec() *NonceSpec {
	return &NonceSpec{
		Nonce: nonce.Nonce,
	}
}
