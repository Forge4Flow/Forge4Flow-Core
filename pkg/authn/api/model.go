package authn

import (
	"time"
)

type Model interface {
	GetID() int64
	GetName() string
	GetKey() *string
	GetExpDate() time.Time
	IsExpired() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
	ToApiSpec() *ApiSpec
}

type ApiKey struct {
	ID          int64      `mysql:"id" postgres:"id" sqlite:"id"`
	DisplayName string     `mysql:"displayName" postgres:"display_name" sqlite:"displayName"`
	ApiKey      string     `mysql:"apikey" postgres:"api_key" sqlite:"apikey"`
	ExpDate     time.Time  `mysql:"expDate" postgres:"exp_date" sqlite:"expDate"`
	CreatedAt   time.Time  `mysql:"createdAt" postgres:"created_at" sqlite:"createdAt"`
	UpdatedAt   time.Time  `mysql:"updatedAt" postgres:"updated_at" sqlite:"updatedAt"`
	DeletedAt   *time.Time `mysql:"deletedAt" postgres:"deleted_at" sqlite:"deletedAt"`
}

func (key ApiKey) GetID() int64 {
	return key.ID
}

func (key ApiKey) GetName() string {
	return key.DisplayName
}

func (key ApiKey) GetKey() *string {
	if key.IsExpired() {
		return nil
	} else {
		return &key.ApiKey
	}
}

func (key ApiKey) GetExpDate() time.Time {
	return key.ExpDate
}

func (key ApiKey) IsExpired() bool {
	return key.ExpDate.Before(time.Now())
}

func (key ApiKey) GetCreatedAt() time.Time {
	return key.CreatedAt
}

func (key ApiKey) GetUpdatedAt() time.Time {
	return key.UpdatedAt
}

func (key ApiKey) GetDeletedAt() *time.Time {
	return key.DeletedAt
}

func (key ApiKey) ToApiSpec() *ApiSpec {
	return &ApiSpec{
		DisplayName: key.DisplayName,
		Key:         key.ApiKey,
		ExpDate:     key.ExpDate,
	}
}
