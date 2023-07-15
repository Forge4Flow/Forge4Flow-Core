package authz

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type PermissionSpec struct {
	PermissionId string    `json:"permissionId" validate:"required,valid_object_id"`
	Name         *string   `json:"name"`
	Description  *string   `json:"description"`
	CreatedAt    time.Time `json:"createdAt"`
}

func (spec PermissionSpec) ToPermission(objectId int64) *Permission {
	return &Permission{
		ObjectId:     objectId,
		PermissionId: spec.PermissionId,
		Name:         spec.Name,
		Description:  spec.Description,
	}
}

func (spec PermissionSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypePermission,
		ObjectId:   spec.PermissionId,
	}
}

type UpdatePermissionSpec struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
