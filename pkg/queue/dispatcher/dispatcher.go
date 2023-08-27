package dispatcher

import "gitub.com/treblle/treblle/pkg/queue"

// Dispatcher creates and controls the workers
type Dispatcher struct {
	// Pool of workers
	WorkerPool chan chan queue.Job
}

func New(maxWorkers int) *Dispatcher {
	pool := make(chan chan queue.Job, maxWorkers)
	return &Dispatcher{WorkerPool: pool}
}

func (d *Dispatcher) Run() {
	for i := 0; i < cap(d.WorkerPool); i++ {
		worker := queue.New(d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for job := range JobQueue {
		go func(job queue.Job) {
			jobChannel := <-d.WorkerPool
			jobChannel <- job
		}(job)
	}
}

var JobQueue = make(chan queue.Job, 100)
