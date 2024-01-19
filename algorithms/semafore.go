package main

import (
	"fmt"
	"sync"
)

func main2() {
	var count int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			count++
			fmt.Println(count)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(count)

}
