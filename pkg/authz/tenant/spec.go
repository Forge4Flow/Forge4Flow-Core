package tenant

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type TenantSpec struct {
	TenantId  string    `json:"tenantId" validate:"omitempty,valid_object_id"`
	Name      *string   `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func (spec TenantSpec) ToTenant(objectId int64) *Tenant {
	return &Tenant{
		ObjectId: objectId,
		TenantId: spec.TenantId,
		Name:     spec.Name,
	}
}

func (spec TenantSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypeTenant,
		ObjectId:   spec.TenantId,
	}
}

type UpdateTenantSpec struct {
	Name *string `json:"name"`
}
