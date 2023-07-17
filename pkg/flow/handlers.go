package flow

import (
	"net/http"

	"github.com/forge4flow/forge4flow-core/pkg/service"
	"github.com/gorilla/websocket"
)

func (svc FlowService) Routes() ([]service.Route, error) {
	return []service.Route{
		service.WarrantRoute{
			Pattern:                    "/v1/flow/events",
			Method:                     "GET",
			Handler:                    service.NewRouteHandler(svc, GetEventsWSHandler),
			OverrideAuthMiddlewareFunc: service.ApiKeyAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/flow/events",
			Method:                     "POST",
			Handler:                    service.NewRouteHandler(svc, AddEventMonitorHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},

		service.WarrantRoute{
			Pattern:                    "/v1/flow/events",
			Method:                     "DELETE",
			Handler:                    service.NewRouteHandler(svc, RemoveEventMonitorHandler),
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
		},
	}, nil
}

func AddEventMonitorHandler(svc FlowService, w http.ResponseWriter, r *http.Request) error {
	var event Event
	err := service.ParseJSONBody(r.Body, &event)
	if err != nil {
		return service.NewInvalidRequestError("Invalid JSON body")
	}

	if event.Type == "" {
		return service.NewMissingRequiredParameterError("Type")
	}

	svc.eventMonitor.AddMonitor(event.Type)

	service.SendJSONResponse(w, event)

	return nil
}

func RemoveEventMonitorHandler(svc FlowService, w http.ResponseWriter, r *http.Request) error {
	var event Event
	err := service.ParseJSONBody(r.Body, &event)
	if err != nil {
		return service.NewInvalidRequestError("Invalid JSON body")
	}

	if event.Type == "" {
		return service.NewMissingRequiredParameterError("Type")
	}

	return svc.eventMonitor.RemoveMonitor(event.Type)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

type SubscriptionRequest struct {
	EventTypes []string `json:"eventTypes"`
}

func GetEventsWSHandler(svc FlowService, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Register the event channel to receive events from the EventMonitorService
	eventChannel := svc.eventMonitor.eventChannel

	// Create a channel to send filtered events to the WebSocket client
	filteredEvents := make(chan Event)

	var subscriptionReq SubscriptionRequest

	// Start a goroutine to filter events based on the client's subscription preferences
	go func() {
		for event := range eventChannel {
			// Check if the event type is in the client's subscribed event types
			if isSubscribed(event.Type, subscriptionReq.EventTypes) {
				filteredEvents <- event
			}
		}
	}()

	// The function will block here until the connection is closed
	for {
		// Read incoming messages from the WebSocket client
		err := conn.ReadJSON(&subscriptionReq)
		if err != nil {
			// Connection closed, so we break the loop
			break
		}
	}

	return nil
}

// Helper function to check if an event type is subscribed by the client
func isSubscribed(eventType string, subscribedEventTypes []string) bool {
	for _, subscribedType := range subscribedEventTypes {
		if subscribedType == eventType {
			return true
		}
	}
	return false
}
