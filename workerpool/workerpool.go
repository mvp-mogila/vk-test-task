package workerpool

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/mvp-mogila/vk-test-task/stack"
)

var ErrNoWorkers = errors.New("worker pool is empty")

type WorkerPool struct {
	wg      *sync.WaitGroup
	lastId  atomic.Int64
	workers *stack.Stack[*Worker]
	taskCh  chan string
}

type Worker struct {
	id       int
	cancelCh chan struct{}
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	wp := WorkerPool{
		wg:      &sync.WaitGroup{},
		workers: stack.NewStack[*Worker](),
		taskCh:  make(chan string),
	}

	for i := 0; i < numWorkers; i++ {
		wp.AddWorker()
	}

	return &wp
}

func (wp *WorkerPool) AddTask(task string) error {
	if wp.workers.Size() != 0 {
		wp.taskCh <- task
		return nil
	}
	return ErrNoWorkers
}

func (wp *WorkerPool) AddWorker() {
	wp.lastId.Add(1)
	worker := &Worker{
		id:       int(wp.lastId.Load()),
		cancelCh: make(chan struct{}),
	}

	wp.workers.Push(worker)
	fmt.Printf("New worker %d added\n", worker.id)

	wp.wg.Add(1)
	go wp.work(worker)
	fmt.Printf("Worker %d started\n", worker.id)
}

func (wp *WorkerPool) RemoveWorker() error {
	worker, err := wp.workers.Pop()
	if err != nil {
		if errors.Is(err, stack.ErrEmptyStack) {
			return ErrNoWorkers
		} else {
			panic(err)
		}
	}

	close(worker.cancelCh)
	fmt.Printf("Worker %d removed\n", worker.id)
	return nil
}

func (wp *WorkerPool) work(w *Worker) {
	defer wp.wg.Done()
	for {
		select {
		case <-w.cancelCh:
			return
		case task, ok := <-wp.taskCh:
			if !ok {
				fmt.Printf("Worker %d stopped\n", w.id)
				return
			}
			fmt.Printf("Worker %d processing string \"%s\"\n", w.id, task)
		}
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.taskCh)
	wp.wg.Wait()
	fmt.Println("Worker pool stopped")
}
