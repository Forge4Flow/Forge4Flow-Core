package flow

import (
	"time"
)

type Model interface {
	GetID() int64
	GetType() string
	GetLastBlockHeight() uint64
	GetObjectType() *string
	GetObjectId() *string
	GetObjectIdField() *string
	GetObjectRelation() *string
	GetSubjectType() *string
	GetSubjectId() *string
	GetSubjectIdField() *string
	GetScript() *string
	GetRemoveAction() bool
	GetActionEnabled() bool
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() *time.Time
	ToEventSpec() *EventSpec
}

type Event struct {
	ID              int64      `mysql:"id" postgres:"id" sqlite:"id"`
	Type            string     `mysql:"type" postgres:"type" sqlite:"type"`
	LastBlockHeight uint64     `mysql:"lastBlockHeight" postgres:"last_block_height" sqlite:"lastBlockHeight"`
	ObjectType      *string    `mysql:"objectType" postgres:"object_type" sqlite:"objectType"`
	ObjectId        *string    `mysql:"objectId" postgres:"object_id" sqlite:"objectId"`
	ObjectIdField   *string    `mysql:"objectIdField" postgres:"object_id_field" sqlite:"objectIdField"`
	ObjectRelation  *string    `mysql:"objectRelation" postgres:"object_relation" sqlite:"objectRelation"`
	SubjectType     *string    `mysql:"subjectType" postgres:"subject_type" sqlite:"subjectType"`
	SubjectId       *string    `mysql:"subjectId" postgres:"subject_id" sqlite:"subjectId"`
	SubjectIdField  *string    `mysql:"subjectIdField" postgres:"subject_id_field" sqlite:"subjectIdField"`
	Script          *string    `mysql:"script" postgres:"script" sqlite:"script"`
	RemoveAction    bool       `mysql:"removeAction" postgres:"remove_action" sqlite:"removeAction"`
	ActionEnabled   bool       `mysql:"actionEnabled" postgres:"action_enabled" sqllite:"actionEnabled"`
	CreatedAt       time.Time  `mysql:"createdAt" postgres:"created_at" sqlite:"createdAt"`
	UpdatedAt       time.Time  `mysql:"updatedAt" postgres:"updated_at" sqlite:"updatedAt"`
	DeletedAt       *time.Time `mysql:"deletedAt" postgres:"deleted_at" sqlite:"deletedAt"`
}

func (event Event) GetID() int64 {
	return event.ID
}

func (event Event) GetType() string {
	return event.Type
}

func (event Event) GetLastBlockHeight() uint64 {
	return event.LastBlockHeight
}

func (event Event) GetObjectType() *string {
	return event.ObjectType
}

func (event Event) GetObjectId() *string {
	return event.ObjectId
}

func (event Event) GetObjectIdField() *string {
	return event.ObjectIdField
}

func (event Event) GetObjectRelation() *string {
	return event.ObjectRelation
}

func (event Event) GetSubjectType() *string {
	return event.SubjectType
}

func (event Event) GetSubjectId() *string {
	return event.SubjectId
}

func (event Event) GetSubjectIdField() *string {
	return event.SubjectIdField
}

func (event Event) GetScript() *string {
	return event.Script
}

func (event Event) GetRemoveAction() bool {
	return event.RemoveAction
}

func (event Event) GetActionEnabled() bool {
	return event.ActionEnabled
}

func (event Event) GetCreatedAt() time.Time {
	return event.CreatedAt
}

func (event Event) GetUpdatedAt() time.Time {
	return event.UpdatedAt
}

func (event Event) GetDeletedAt() *time.Time {
	return event.DeletedAt
}

func (event Event) ToEventSpec() *EventSpec {
	return &EventSpec{
		Type:           event.Type,
		ObjectType:     *event.ObjectType,
		ObjectId:       *event.ObjectId,
		ObjectIdField:  *event.ObjectIdField,
		ObjectRelation: *event.ObjectRelation,
		SubjectType:    *event.SubjectType,
		SubjectId:      *event.SubjectId,
		SubjectIdField: *event.SubjectIdField,
		Script:         *event.Script,
		RemoveAction:   event.RemoveAction,
		ActionEnabled:  event.ActionEnabled,
	}
}
