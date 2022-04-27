package worker

import "sync"

// worker is a struct that contains input channel of tasks
type worker struct {
	tasks     chan Task
	quit      chan struct{}
	isRunning bool
	mut       sync.Mutex
}

// NewWorker creates a new worker with input channel of tasks
// and returns a pointer to the worker
// The worker is not started yet
func NewWorker(tasks chan Task) *worker {
	return &worker{
		tasks: tasks,
		quit:  make(chan struct{}),
	}
}

// Start starts the worker loop to consume and run tasks
func (w *worker) Start() {
	w.mut.Lock()
	if w.isRunning {
		w.mut.Unlock()
		return
	}
	w.isRunning = true
	w.mut.Unlock()

	go func() {
		for {
			select {
			case job, ok := <-w.tasks:
				if !ok {
					w.mut.Lock()
					defer w.mut.Unlock()
					w.isRunning = false
					return
				}
				job.Run()
			case <-w.quit:
				return
			}
		}
	}()
}

// Stop stops the worker loop
func (w *worker) Stop() {
	w.mut.Lock()
	defer w.mut.Unlock()
	if !w.isRunning {
		return
	}
	w.quit <- struct{}{}
	w.isRunning = false
}

// IsRunning returns true if the worker is running
func (w *worker) IsRunning() bool {
	w.mut.Lock()
	defer w.mut.Unlock()
	return w.isRunning
}

// language: go
