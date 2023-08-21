package flow

type EventActionsSpec struct {
	Type           string `json:"type,omitempty"`
	ObjectType     string `json:"objectType,omitempty" validate:"required_with=ActionEnabled"`
	ObjectId       string `json:"objectId,omitempty"`
	ObjectIdField  string `json:"objectIdField,omitempty"`
	ObjectRelation string `json:"objectRelation,omitempty"`
	SubjectType    string `json:"subjectType,omitempty"`
	SubjectId      string `json:"subjectId"`
	SubjectIdField string `json:"subjectIdField"`
	Script         string `json:"script,omitempty"`
	OrderWeight    int64  `json:"orderWeight" validate:"required"`
	RemoveAction   bool   `json:"removeAction,omitempty"`
	ActionEnabled  bool   `json:"actionEnabled,omitempty"`
}

func (e *EventActionsSpec) ToEventAction() EventAction {
	return EventAction{
		Type:           e.Type,
		ObjectType:     &e.ObjectType,
		ObjectId:       &e.ObjectId,
		ObjectIdField:  &e.ObjectIdField,
		ObjectRelation: &e.ObjectRelation,
		SubjectType:    &e.SubjectType,
		SubjectId:      &e.SubjectId,
		SubjectIdField: &e.SubjectIdField,
		Script:         &e.Script,
		OrderWeight:    e.OrderWeight,
		RemoveAction:   e.RemoveAction,
		ActionEnabled:  e.ActionEnabled,
	}
}
