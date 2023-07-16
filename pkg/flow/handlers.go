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
			OverrideAuthMiddlewareFunc: service.PassthroughAuthMiddleware,
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
			Handler:                    service.NewRouteHandler(svc, GetEventsWSHandler),
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

	svc.eventMonitor.AddMonitor(event.Type)

	service.SendJSONResponse(w, event)

	return nil
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func GetEventsWSHandler(svc FlowService, w http.ResponseWriter, r *http.Request) error {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	defer conn.Close()

	// Register the event channel to receive events from the EventMonitorService
	eventChannel := svc.eventMonitor.eventChannel

	// Start a goroutine to send events from the event channel to the WebSocket client
	go func() {
		for event := range eventChannel {
			err := conn.WriteJSON(event)
			if err != nil {
				// Handle error, e.g., log and break the loop
				break
			}
		}
	}()

	// The function will block here until the connection is closed
	for {
		// Read incoming messages from the WebSocket client
		_, _, err := conn.ReadMessage()
		if err != nil {
			// Connection closed, so we break the loop
			break
		}
	}

	return nil
}
