package flow

type EventSpec struct {
	Type          string      `json:"type,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	TransactionID string      `json:"transaction_id,omitempty"`
}

func (e *EventSpec) ToEvent() Event {
	return Event{
		Type: e.Type,
	}
}
