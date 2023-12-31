package authz

import (
	"context"

	objecttype "github.com/forge4flow/forge4flow-core/pkg/authz/objecttype"
	"github.com/forge4flow/forge4flow-core/pkg/event"
	"github.com/forge4flow/forge4flow-core/pkg/service"
)

type WarrantService struct {
	service.BaseService
	Repository    WarrantRepository
	EventSvc      *event.EventService
	ObjectTypeSvc *objecttype.ObjectTypeService
}

func NewService(env service.Env, repository WarrantRepository, eventSvc *event.EventService, objectTypeSvc *objecttype.ObjectTypeService) *WarrantService {
	return &WarrantService{
		BaseService:   service.NewBaseService(env),
		Repository:    repository,
		EventSvc:      eventSvc,
		ObjectTypeSvc: objectTypeSvc,
	}
}

func (svc WarrantService) ID() string {
	return service.WarrantService
}

func (svc WarrantService) Create(ctx context.Context, warrantSpec WarrantSpec) (*WarrantSpec, error) {
	// Check that objectType is valid
	objectTypeDef, err := svc.ObjectTypeSvc.GetByTypeId(ctx, warrantSpec.ObjectType)
	if err != nil {
		return nil, service.NewInvalidParameterError("objectType", "The given object type does not exist.")
	}

	// Check that relation is valid for objectType
	_, exists := objectTypeDef.Relations[warrantSpec.Relation]
	if !exists {
		return nil, service.NewInvalidParameterError("relation", "An object type with the given relation does not exist.")
	}

	// Check that warrant does not already exist
	_, err = svc.Repository.Get(ctx, warrantSpec.ObjectType, warrantSpec.ObjectId, warrantSpec.Relation, warrantSpec.Subject.ObjectType, warrantSpec.Subject.ObjectId, warrantSpec.Subject.Relation, warrantSpec.Policy.Hash())
	if err == nil {
		return nil, service.NewDuplicateRecordError("Warrant", warrantSpec, "A warrant with the given objectType, objectId, relation, subject, and policy already exists")
	}

	var createdWarrant Model
	err = svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		warrant, err := warrantSpec.ToWarrant()
		if err != nil {
			return err
		}

		createdWarrantId, err := svc.Repository.Create(txCtx, warrant)
		if err != nil {
			return err
		}

		createdWarrant, err = svc.Repository.GetByID(txCtx, createdWarrantId)
		if err != nil {
			return err
		}

		var eventMeta map[string]interface{}
		if createdWarrant.GetPolicy() != "" {
			eventMeta = make(map[string]interface{})
			eventMeta["policy"] = createdWarrant.GetPolicy()
		}

		err = svc.EventSvc.TrackAccessGrantedEvent(txCtx, createdWarrant.GetObjectType(), createdWarrant.GetObjectId(), createdWarrant.GetRelation(), createdWarrant.GetSubjectType(), createdWarrant.GetSubjectId(), createdWarrant.GetSubjectRelation(), eventMeta)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return createdWarrant.ToWarrantSpec(), nil
}

func (svc WarrantService) List(ctx context.Context, filterOptions *FilterOptions, listParams service.ListParams) ([]*WarrantSpec, error) {
	warrantSpecs := make([]*WarrantSpec, 0)
	warrants, err := svc.Repository.List(ctx, filterOptions, listParams)
	if err != nil {
		return nil, err
	}

	for _, warrant := range warrants {
		warrantSpecs = append(warrantSpecs, warrant.ToWarrantSpec())
	}

	return warrantSpecs, nil
}

func (svc WarrantService) Delete(ctx context.Context, warrantSpec WarrantSpec) error {
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		warrantToDelete, err := warrantSpec.ToWarrant()
		if err != nil {
			return nil
		}

		warrant, err := svc.Repository.Get(txCtx, warrantToDelete.GetObjectType(), warrantToDelete.GetObjectId(), warrantToDelete.GetRelation(), warrantToDelete.GetSubjectType(), warrantToDelete.GetSubjectId(), warrantToDelete.GetSubjectRelation(), warrantToDelete.GetPolicyHash())
		if err != nil {
			return err
		}

		err = svc.Repository.DeleteById(txCtx, warrant.GetID())
		if err != nil {
			return err
		}

		var eventMeta map[string]interface{}
		if warrantSpec.Policy != "" {
			eventMeta = make(map[string]interface{})
			eventMeta["policy"] = warrantSpec.Policy
		}

		err = svc.EventSvc.TrackAccessRevokedEvent(txCtx, warrantSpec.ObjectType, warrantSpec.ObjectId, warrantSpec.Relation, warrantSpec.Subject.ObjectType, warrantSpec.Subject.ObjectId, warrantSpec.Subject.Relation, eventMeta)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc WarrantService) DeleteRelatedWarrants(ctx context.Context, objectType string, objectId string) error {
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		err := svc.Repository.DeleteAllByObject(txCtx, objectType, objectId)
		if err != nil {
			return err
		}

		err = svc.Repository.DeleteAllBySubject(txCtx, objectType, objectId)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (svc WarrantService) GetAllWarrantsForSubjectIdByType(ctx context.Context, subject string, subjectId string, objectType string) ([]Model, error) {
	matchingWarrants := make([]Model, 0)
	// Get all warrants where subjectType is subject and subjectid is subjectId
	directWarrants, err := svc.Repository.GetAllMatchingSubjectId(ctx, subject, subjectId)
	if err != nil {
		return matchingWarrants, nil
	}

	for _, directWarrant := range directWarrants {
		svc.processWarrant(ctx, directWarrant, objectType, &matchingWarrants)
	}

	return matchingWarrants, nil
}

func (svc WarrantService) processWarrant(ctx context.Context, warrantObject Model, objectType string, matchingWarrants *[]Model) {
	// if objectTypes don't match
	if warrantObject.GetObjectType() != objectType {
		indirectWarrants := make([]Model, 0)
		var err error
		// Does warrant have a subjectRelation
		if warrantObject.GetSubjectRelation() != "" {
			// get all warrants where subjectType, subjectId, and subjectRelationship match current warrant objectType & objectId
			indirectWarrants, err = svc.Repository.GetAllMatchingSubject(ctx, warrantObject.GetObjectType(), warrantObject.GetObjectId(), warrantObject.GetRelation())
			if err != nil {
				return
			}
		} else {
			indirectWarrants, err = svc.Repository.GetAllMatchingSubjectId(ctx, warrantObject.GetObjectType(), warrantObject.GetObjectId())
			if err != nil {
				return
			}
		}

		// Check if indirect warrants ObjectType matches requested objectType
		for _, indirectWarrant := range indirectWarrants {
			if indirectWarrant.GetObjectType() != objectType {
				// repeat until warrant objectType matches requested objectType
				svc.processWarrant(ctx, indirectWarrant, objectType, matchingWarrants)
			} else {
				*matchingWarrants = append(*matchingWarrants, indirectWarrant)
			}
		}

		return
	}

	// If warrant objectType matches requested objectType add to matchingWarrants
	*matchingWarrants = append(*matchingWarrants, warrantObject)
}
