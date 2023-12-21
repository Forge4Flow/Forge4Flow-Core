package authn

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type ApiSpec struct {
	DisplayName string    `json:"displayName" validate:"required"`
	ApiKey      string    `json:"apiKey" validate:"omitempty"`
	ExpDate     time.Time `json:"expDate" validate:"omitempty"`
}

func (spec ApiSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypeApiKey,
		ObjectId:   spec.ApiKey,
	}
}

func (spec ApiSpec) ToApiKey(objectId int64) *ApiKey {
	return &ApiKey{
		ObjectId:    objectId,
		DisplayName: spec.DisplayName,
		ApiKey:      spec.ApiKey,
		ExpDate:     spec.ExpDate,
	}
}
