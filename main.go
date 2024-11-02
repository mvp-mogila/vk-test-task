package main

import (
	"time"

	"github.com/mvp-mogila/vk-test-task/workerpool"
)

func main() {
	wp := workerpool.NewWorkerPool(2)

	strings := []string{"123", "abc", "cde", "fhurekf", "fherujkferf", "fjhureofref", "hjfuerofer",
		"123", "abc", "cde", "fhurekf", "gtr", "fjhureofref", "hjfuerofer",
		"45", "i", "cde", "op", "fherujkferf", "fjhureofref", "hjfuerofer",
		"65", "abc", "lio", "fhurekf", "grt", "grt", "hjfuerofer"}

	for i, s := range strings {
		wp.AddTask(s)
		if i == 2 {
			wp.RemoveWorker()
		} else if i == 4 {
			wp.AddWorker()
		}
		wp.AddWorker()
	}

	time.Sleep(2 * time.Second)
	wp.Stop()
}
