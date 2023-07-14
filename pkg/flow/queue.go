package flow

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type JobInterface interface {
	Execute()
}

type JobWithID struct {
	Job JobInterface
	ID  string
}

type Worker struct {
	ID         int
	JobChannel chan JobWithID
}

type Queue struct {
	Workers      []Worker
	RateLimit    int
	JobChannel   chan JobWithID
	WaitGroup    sync.WaitGroup
	LimiterMutex sync.Mutex
	stopChannel  chan struct{}
	Jobs         map[string]bool
	JobsMutex    sync.Mutex
	LastJobID    int
	flowSvc      *FlowService
	running      bool
}

func newQueue(svc *FlowService, rateLimit int) *Queue {
	return &Queue{
		RateLimit:   rateLimit,
		JobChannel:  make(chan JobWithID),
		stopChannel: make(chan struct{}),
		Jobs:        make(map[string]bool),
		LastJobID:   0,
		flowSvc:     svc,
	}
}

func (q *Queue) Start() {
	if q.running {
		return
	}
	q.running = true

	numWorkers := q.RateLimit
	q.Workers = make([]Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		q.Workers[i] = Worker{
			ID:         i + 1,
			JobChannel: make(chan JobWithID),
		}
		go q.worker(i + 1)
	}

	ticker := time.NewTicker(time.Second / time.Duration(q.RateLimit))

	for {
		select {
		case <-q.stopChannel:
			ticker.Stop()
			q.running = false
			return
		case <-ticker.C:
			select {
			case job := <-q.JobChannel:
				<-q.Workers[0].JobChannel
				q.Workers[0].JobChannel <- job
			default:
				// No job in the JobChannel, do nothing
			}
		}
	}
}

func (q *Queue) worker(id int) {
	for jobWithID := range q.Workers[id-1].JobChannel {
		select {
		case <-q.stopChannel:
			return
		default:
			q.WaitGroup.Add(1)
			q.processJob(jobWithID)
			q.WaitGroup.Done()
			q.JobsMutex.Lock()
			delete(q.Jobs, jobWithID.ID)
			q.JobsMutex.Unlock()
		}
	}
}

func (q *Queue) processJob(jobWithId JobWithID) {
	fmt.Printf("Processing job %s:", jobWithId.ID)
	jobWithId.Job.Execute()
	fmt.Printf("Job %s completed\n", jobWithId.ID)
}

func (q *Queue) Stop() {
	if q.running {
		close(q.stopChannel)
	}
}

func (q *Queue) RemoveJobByID(id string) error {
	q.JobsMutex.Lock()
	_, ok := q.Jobs[id]
	if !ok {
		return errors.New("EventMonitor not found: " + id)
	}
	delete(q.Jobs, id)
	q.JobsMutex.Unlock()

	return nil
}

func (q *Queue) CreateJob(job JobInterface) (string, error) {
	q.JobsMutex.Lock()
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	jobID := newUUID.String()
	q.Jobs[jobID] = true
	q.JobsMutex.Unlock()

	jobWithID := JobWithID{
		Job: job,
		ID:  jobID,
	}

	go func() {
		q.JobChannel <- jobWithID
	}()

	return jobID, nil
}
