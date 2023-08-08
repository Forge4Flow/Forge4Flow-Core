package flow

type EventSpec struct {
	Type           string      `json:"type,omitempty" validate:"required"`
	ObjectType     string      `json:"objectType,omitempty" validate:"required_with=ActionEnabled"`
	ObjectId       string      `json:"objectId,omitempty"`
	ObjectIdField  string      `json:"objectIdField,omitempty"`
	ObjectRelation string      `json:"objectRelation,omitempty"`
	SubjectType    string      `json:"subjectType,omitempty"`
	SubjectId      string      `json:"subjectId"`
	SubjectIdField string      `json:"subjectIdField"`
	Script         string      `json:"script,omitempty"`
	RemoveAction   bool        `json:"removeAction,omitempty"`
	ActionEnabled  bool        `json:"actionEnabled,omitempty"`
	Data           interface{} `json:"data,omitempty"`
	TransactionID  string      `json:"transaction_id,omitempty"`
}

func (e *EventSpec) ToEvent() Event {
	return Event{
		Type:           e.Type,
		ObjectType:     &e.ObjectType,
		ObjectId:       &e.ObjectId,
		ObjectIdField:  &e.ObjectIdField,
		ObjectRelation: &e.ObjectRelation,
		SubjectType:    &e.SubjectType,
		SubjectId:      &e.SubjectId,
		SubjectIdField: &e.SubjectIdField,
		Script:         &e.Script,
		RemoveAction:   e.RemoveAction,
		ActionEnabled:  e.ActionEnabled,
	}
}
