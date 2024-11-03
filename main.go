package main

import (
	"fmt"
	"sync"

	"github.com/mvp-mogila/vk-test-task/workerpool"
)

func main() {
	wp := workerpool.NewWorkerPool(1)
	wp.RemoveWorker()              // id: 1
	fmt.Println(wp.RemoveWorker()) // error

	fmt.Println(wp.AddTask("test")) // error

	wp.AddWorker() // id: 2
	wp.AddWorker() // id: 3
	wp.AddWorker() // id: 4

	strings := []string{"123", "abc", "cde", "fhurekf", "fherujkferf", "fjhurevfvv.ofref", "hjfuerofer",
		"12343", "abzzc", "cxde", "fzchurekf", "gtr", "fjhureofref", "5",
		"45", "i", "cdcccce", "op", "fherujkferf", "fjhur453eofref", "hjfuerofer",
		"65", "abc", "lio", "fhuvfrekf", "grt", "grtf54", "hjfuerozxczfer",
		"12ghyth343", "ab546456zzc", "cxdhyte", "fzchunnnnnrekf", "gn)(*(*tr", "fjhur!!!!!!!!eofref", ""}

	wg := sync.WaitGroup{}
	for i, s := range strings {
		if i == 4 {
			wp.AddWorker() // id: 5
		} else if i == 9 {
			wp.RemoveWorker() // id: 5
		}

		wg.Add(1)
		go func() {
			wp.AddTask(s)
			wg.Done()
		}()
	}

	wp.RemoveWorker() // id: 4

	wg.Wait()

	wp.Stop()
}
