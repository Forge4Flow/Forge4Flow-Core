package flow

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/forge4flow/forge4flow-core/pkg/service"
)

func (svc *FlowService) AddEventMonitor(ctx context.Context, event EventSpec) error {
	// Add to DB
	var newEvent Model
	err := svc.Env().DB().WithinTransaction(ctx, func(txCtx context.Context) error {
		_, err := svc.Repository.GetByType(txCtx, event.Type)
		if err == nil {
			return service.NewDuplicateRecordError("FlowEventType", event.Type, "A event monitor with the given event type already exists")
		}

		newEventId, err := svc.Repository.Create(txCtx, event.ToEvent())
		if err != nil {
			return err
		}

		newEvent, err = svc.Repository.GetById(txCtx, newEventId)
		if err != nil {
			return err
		}

		// TODO: Integrate with EventSvc for monitoring and logging
		// err = svc.EventSvc.TrackResourceCreated(txCtx, ResourceTypePermission, newPermission.GetPermissionId(), newPermission.ToPermissionSpec())
		// if err != nil {
		// 	return err
		// }

		return nil
	})
	if err != nil {
		return err
	}

	// Add to EventMonitor Service
	svc.eventMonitor.AddMonitor(newEvent.GetType())

	return nil
}

type EventMonitorService struct {
	monitors      map[string]*EventMonitor
	monitorsMutex sync.Mutex
	flowSvc       *FlowService
	eventChannel  chan EventSpec // Channel to receive events
}

func newEventMonitorService(svc *FlowService) *EventMonitorService {
	return &EventMonitorService{
		flowSvc:      svc,
		monitors:     make(map[string]*EventMonitor),
		eventChannel: make(chan EventSpec),
	}
}

func (ems *EventMonitorService) StartService() {
	ems.monitorsMutex.Lock()
	defer ems.monitorsMutex.Unlock()

	for _, em := range ems.monitors {
		em.Start()
	}
}

func (ems *EventMonitorService) StopService() {
	ems.monitorsMutex.Lock()
	defer ems.monitorsMutex.Unlock()

	for _, em := range ems.monitors {
		em.Stop()
	}
}

func (ems *EventMonitorService) AddMonitor(eventID string) {
	em := &EventMonitor{
		EventID:      eventID,
		stopChan:     make(chan struct{}),
		flowSvc:      ems.flowSvc,
		queue:        ems.flowSvc.queue,
		eventChannel: ems.eventChannel,
	}

	ems.monitorsMutex.Lock()
	ems.monitors[eventID] = em
	ems.monitorsMutex.Unlock()

	if ems.flowSvc.queue != nil {
		em.Start()
	}
}

func (ems *EventMonitorService) RemoveMonitor(eventID string) error {
	ems.monitorsMutex.Lock()
	defer ems.monitorsMutex.Unlock()

	_, ok := ems.monitors[eventID]
	if !ok {
		return errors.New("EventMonitor not found: " + eventID)
	}

	delete(ems.monitors, eventID)

	return nil
}

func (ems *EventMonitorService) StartMonitor(eventID string) error {
	ems.monitorsMutex.Lock()
	defer ems.monitorsMutex.Unlock()

	em, ok := ems.monitors[eventID]
	if !ok {
		return errors.New("EventMonitor not found: " + eventID)
	}

	if ems.flowSvc.queue != nil {
		em.Start()
	}

	return nil
}

func (ems *EventMonitorService) StopMonitor(eventID string) error {
	ems.monitorsMutex.Lock()
	defer ems.monitorsMutex.Unlock()

	em, ok := ems.monitors[eventID]
	if !ok {
		return errors.New("EventMonitor not found: " + eventID)
	}

	if ems.flowSvc.queue != nil {
		em.Stop()
	}

	return nil
}

type Job struct {
	ExecuteFunc func()
	Done        chan struct{} // Channel to signal job completion
}

func (j *Job) Execute() {
	fmt.Println("executing called")
	if j.ExecuteFunc != nil {
		j.ExecuteFunc()
	}
}

type EventMonitor struct {
	EventID         string `json:"eventId"`
	lastBlockHeight uint64
	running         bool
	stopChan        chan struct{}
	mutex           sync.Mutex
	flowSvc         *FlowService
	queue           *Queue
	eventChannel    chan<- EventSpec // Channel to send events to the service
}

func (em *EventMonitor) Start() {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if em.running {
		return
	}

	go em.runLoop()
}

func (em *EventMonitor) Stop() {
	em.mutex.Lock()
	defer em.mutex.Unlock()

	if !em.running {
		return
	}

	close(em.stopChan)
}

func (em *EventMonitor) runLoop() {
	fmt.Println("is running 1")
	em.running = true
	for {
		select {
		case <-em.stopChan:
			fmt.Println("stop called")
			em.running = false
			return
		default:
			fmt.Println("is running loop")
			// Create a new job to be executed by the queue
			job := &Job{
				Done: make(chan struct{}),
			}

			job.ExecuteFunc = func() {
				fmt.Println("job execution")
				ctx := context.Background()

				// Get last sealed block height from blockchain
				latestBlock, err := em.flowSvc.FlowClient.GetLatestBlock(ctx, true)
				if err != nil {
					log.Println(err)
				}

				if em.lastBlockHeight == 0 {
					em.lastBlockHeight = latestBlock.Height
				}

				if latestBlock.Height-em.lastBlockHeight > 250 {
					latestBlock.Height = em.lastBlockHeight + 200
				}

				// Query events from block range
				blocks, err := em.flowSvc.FlowClient.GetEventsForHeightRange(ctx, em.EventID, em.lastBlockHeight, latestBlock.Height)
				if err != nil {
					log.Println(err)
				}

				// Updated last block height
				em.lastBlockHeight = latestBlock.Height
				em.flowSvc.Repository.UpdateLastBlockHeightByType(context.Background(), em.EventID, latestBlock.Height)

				// parse events from block range
				for _, block := range blocks {
					for _, cadenceEvent := range block.Events {
						log.Printf("\n\nType: %s", cadenceEvent.Type)
						log.Printf("\nValues: %v", cadenceEvent.Value)
						log.Printf("\nTransaction ID: %s", cadenceEvent.TransactionID)

						event := EventSpec{
							Type:          cadenceEvent.Type,
							Data:          CadenceValueToInterface(cadenceEvent.Value),
							TransactionID: cadenceEvent.TransactionID.Hex(),
						}

						em.eventChannel <- event
					}
				}
			}

			em.queue.CreateJob(job)

			// Wait for the job to complete
			time.Sleep(5 * time.Second)
			fmt.Println("job done")
		}
	}
}
