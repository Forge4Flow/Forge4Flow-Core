package flow

type EventSpec struct {
	Type          string      `json:"type,omitempty"`
	ObjectType    string      `json:"objectType,omitempty"`
	ObjectIdField string      `json:"objectIdField,omitempty"`
	OwnerField    string      `json:"ownerField,omitempty"`
	Script        string      `json:"script,omitempty"`
	RemovedAction bool        `json:"removedAction,omitempty"`
	ActionEnabled bool        `json:"actionEnabled,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	TransactionID string      `json:"transaction_id,omitempty"`
}

func (e *EventSpec) ToEvent() Event {
	return Event{
		Type: e.Type,
		// ObjectType:    &e.ObjectType,
		// ObjectIdField: &e.ObjectIdField,
		// OwnerField:    &e.OwnerField,
		// Script:        &e.Script,
		// RemoveAction:  &e.RemovedAction,
		// ActionEnabled: e.ActionEnabled,
	}
}
