package pworker

import "sync"

type workerPool struct {
	mux         sync.Mutex
	taskChannel chan Task
	workers     []*worker
}

// NewWorkerPool creates a new worker workerPool
// with the given number of workers and the given task channel.
// The task channel is used to send tasks to the workers.
// The workers are started immediately.
// The workers are stopped when the workerPool is stopped.
func NewWorkerPool(taskChannel chan Task, numWorkers int) *workerPool {
	pool := &workerPool{
		workers:     make([]*worker, numWorkers),
		taskChannel: taskChannel,
	}

	for i := 0; i < numWorkers; i++ {
		w := NewWorker(taskChannel)
		pool.workers[i] = w
		w.Start()
	}

	return pool
}

// Stop stops all workers in the workerPool.
func (p *workerPool) Stop() {
	p.mux.Lock()
	defer p.mux.Unlock()
	for i := 0; i <= len(p.workers); i++ {
		if len(p.workers) > 0 {
			w := p.workers[len(p.workers)-1]
			w.Stop()
			p.workers = p.workers[:len(p.workers)-1]
		}
	}
}

// Add a single task to process queue
func (p *workerPool) AddTask(task Task) {
	p.taskChannel <- task
}

// Reurns the queue chanf of the workerPool
func (p *workerPool) TaskChannel() chan Task {
	return p.taskChannel
}

// Returns the number of workers in the workerPool
func (p *workerPool) WorkersCount() int {
	return len(p.workers)
}

// AddWorkers adds the given number of workers to the workerPool.
func (p *workerPool) AddWorkers(count int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	for i := 0; i < count; i++ {
		w := NewWorker(p.taskChannel)
		p.workers = append(p.workers, w)
		w.Start()
	}
}

// RemoveWorkers removes the given number of workers from the workerPool.
// The workers are stopped and removed from the workerPool.
// If the number of workers to remove is greater than the number of workers in the workerPool,
// all workers are removed.
func (p *workerPool) RemoveWorkers(count int) {
	p.mux.Lock()
	defer p.mux.Unlock()
	for i := 0; i < count; i++ {
		if len(p.workers) > 0 {
			w := p.workers[len(p.workers)-1]
			w.Stop()
			p.workers = p.workers[:len(p.workers)-1]
		}
	}
}

// Language: go
