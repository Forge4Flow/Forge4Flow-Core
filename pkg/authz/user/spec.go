package authz

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type UserSpec struct {
	UserId    string    `json:"userId" validate:"omitempty,valid_object_id"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

func (spec UserSpec) ToUser(objectId int64) *User {
	return &User{
		ObjectId: objectId,
		UserId:   spec.UserId,
		Email:    spec.Email,
	}
}

func (spec UserSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypeUser,
		ObjectId:   spec.UserId,
	}
}

type UpdateUserSpec struct {
	Email *string `json:"email"`
}
