package main

import (
	"fmt"
	"sync"
)

func taskId(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("Running task id", id)
}

func main() {

	var wg sync.WaitGroup

	for i := range 10 {
		wg.Add(1)

		go func (i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)

		// go taskId(i, &wg)
	}

	// time.Sleep(time.Second * 1)                 // use Wait groups instead
	wg.Wait()
}
