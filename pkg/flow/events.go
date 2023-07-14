package flow

import (
	"context"
	"errors"
	"log"
	"sync"
)

type EventMonitorService struct {
	monitors      map[string]*EventMonitor
	monitorsMutex sync.Mutex
	flowSvc       *FlowService
}

func newEventMonitorService(svc *FlowService) *EventMonitorService {
	return &EventMonitorService{
		flowSvc: svc,
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
		EventID:  eventID,
		stopChan: make(chan struct{}),
		flowSvc:  ems.flowSvc,
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
	em.running = true
	for {
		select {
		case <-em.stopChan:
			em.running = false
			return
		default:
			// Create a new job to be executed by the queue
			job := &Job{
				Done: make(chan struct{}),
			}

			// Capture the `job` variable in a closure
			go func(job *Job) {
				job.ExecuteFunc = func() {
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
						for _, event := range block.Events {
							log.Printf("\n\nType: %s", event.Type)
							log.Printf("\nValues: %v", event.Value)
							log.Printf("\nTransaction ID: %s", event.TransactionID)
						}
					}

					// Signal job completion
					close(job.Done)
				}
			}(job)

			go em.queue.CreateJob(job)

			// Wait for the job to complete
			<-job.Done
		}
	}
}
