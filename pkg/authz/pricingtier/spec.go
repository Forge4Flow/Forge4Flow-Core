package authz

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type PricingTierSpec struct {
	PricingTierId string    `json:"pricingTierId" validate:"required,valid_object_id"`
	Name          *string   `json:"name"`
	Description   *string   `json:"description"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (spec PricingTierSpec) ToPricingTier(objectId int64) Model {
	return &PricingTier{
		ObjectId:      objectId,
		PricingTierId: spec.PricingTierId,
		Name:          spec.Name,
		Description:   spec.Description,
	}
}

func (spec PricingTierSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypePricingTier,
		ObjectId:   spec.PricingTierId,
	}
}

type UpdatePricingTierSpec struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
