package workerpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	wg      *sync.WaitGroup
	mu      *sync.Mutex
	lastId  int
	workers []*Worker
	taskCh  chan string
}

type Worker struct {
	id       int
	cancelCh chan struct{}
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	wp := WorkerPool{
		wg:      &sync.WaitGroup{},
		mu:      &sync.Mutex{},
		workers: make([]*Worker, 0, numWorkers),
		taskCh:  make(chan string),
	}

	for i := 0; i < numWorkers; i++ {
		wp.AddWorker()
	}

	return &wp
}

func (wp *WorkerPool) AddTask(task string) {
	wp.taskCh <- task
}

func (wp *WorkerPool) AddWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	wp.lastId++
	worker := &Worker{
		id:       wp.lastId,
		cancelCh: make(chan struct{}),
	}

	//TODO: use allocated memory after removing worker
	wp.workers = append(wp.workers, worker)
	fmt.Printf("New worker %d added\n", worker.id)

	wp.wg.Add(1)
	go wp.worker(worker)
	fmt.Printf("Worker %d started\n", worker.id)
}

func (wp *WorkerPool) RemoveWorker() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	worker := wp.workers[wp.lastId-1]
	worker.cancelCh <- struct{}{}
	wp.workers = wp.workers[:wp.lastId-1]
	//wp.numWorkers--

	fmt.Printf("Worker %d removed\n", worker.id)
}

func (wp *WorkerPool) worker(w *Worker) {
	defer wp.wg.Done()
	for {
		select {
		case task, ok := <-wp.taskCh:
			if !ok {
				fmt.Printf("Worker %d stopped\n", w.id)
				return
			}
			fmt.Printf("Worker %d processing string %s\n", w.id, task)
		case <-w.cancelCh:
			fmt.Printf("Worker %d stopped\n", w.id)
			return
		}
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.taskCh)
	wp.wg.Wait()
}
