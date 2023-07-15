package authz

import (
	"time"

	object "github.com/forge4flow/forge4flow-core/pkg/authz/object"
	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
)

type FeatureSpec struct {
	FeatureId   string    `json:"featureId" validate:"required,valid_object_id"`
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

func (spec FeatureSpec) ToFeature(objectId int64) *Feature {
	return &Feature{
		ObjectId:    objectId,
		FeatureId:   spec.FeatureId,
		Name:        spec.Name,
		Description: spec.Description,
	}
}

func (spec FeatureSpec) ToObjectSpec() *object.ObjectSpec {
	return &object.ObjectSpec{
		ObjectType: objecttype.ObjectTypeFeature,
		ObjectId:   spec.FeatureId,
	}
}

type UpdateFeatureSpec struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}
