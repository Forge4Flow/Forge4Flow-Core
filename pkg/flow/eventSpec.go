package flow

type EventSpec struct {
	Type          string      `json:"type,omitempty" validate:"required"`
	ObjectType    string      `json:"objectType,omitempty" validate:"required_with=ActionEnabled"`
	ObjectIdField string      `json:"objectIdField,omitempty" validate:"required_with=ActionEnabled"`
	OwnerField    string      `json:"ownerField,omitempty" validate:"required_with=ActionEnabled"`
	Script        string      `json:"script,omitempty" validate:"required_with=ActionEnabled"`
	RemovedAction bool        `json:"removedAction,omitempty" validate:"required_with=ActionEnabled"`
	ActionEnabled bool        `json:"actionEnabled,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	TransactionID string      `json:"transaction_id,omitempty"`
}

func (e *EventSpec) ToEvent() Event {
	return Event{
		Type:          e.Type,
		ObjectType:    &e.ObjectType,
		ObjectIdField: &e.ObjectIdField,
		OwnerField:    &e.OwnerField,
		Script:        &e.Script,
		RemoveAction:  e.RemovedAction,
		ActionEnabled: e.ActionEnabled,
	}
}
