package flow

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
)

type Event struct {
	Type          string      `json:"type,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	TransactionID string      `json:"transaction_id,omitempty"`
}

type EventMonitorService struct {
	monitors      map[string]*EventMonitor
	monitorsMutex sync.Mutex
	flowSvc       *FlowService
	eventChannel  chan Event // Channel to receive events
}

func newEventMonitorService(svc *FlowService) *EventMonitorService {
	return &EventMonitorService{
		flowSvc:      svc,
		monitors:     make(map[string]*EventMonitor),
		eventChannel: make(chan Event),
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
		close(j.Done)
	}
}

func (j *Job) Close() {
	close(j.Done)
}

type EventMonitor struct {
	EventID         string `json:"eventId"`
	lastBlockHeight uint64
	running         bool
	stopChan        chan struct{}
	mutex           sync.Mutex
	flowSvc         *FlowService
	queue           *Queue
	eventChannel    chan<- Event // Channel to send events to the service
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

				// Query events from block range
				blocks, err := em.flowSvc.FlowClient.GetEventsForHeightRange(ctx, em.EventID, em.lastBlockHeight, latestBlock.Height)
				if err != nil {
					log.Println(err)
				}

				// Updated last block height
				em.lastBlockHeight = latestBlock.Height

				// parse events from block range
				for _, block := range blocks {
					for _, cadenceEvent := range block.Events {
						log.Printf("\n\nType: %s", cadenceEvent.Type)
						log.Printf("\nValues: %v", cadenceEvent.Value)
						log.Printf("\nTransaction ID: %s", cadenceEvent.TransactionID)

						event := Event{
							Type:          cadenceEvent.Type,
							Data:          CadenceValueToInterface(cadenceEvent.Value),
							TransactionID: cadenceEvent.TransactionID.Hex(),
						}

						em.eventChannel <- event
					}
				}
				close(job.Done)
			}

			em.queue.CreateJob(job)

			// Wait for the job to complete
			<-job.Done
			fmt.Println("job done")
		}
	}
}
