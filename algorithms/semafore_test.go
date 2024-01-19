package algorithms

import (
	"sync"
	"testing"
)

func Test_main2(t *testing.T) {
	var count int = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(10000)

	for i := 0; i < 10000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			count++
			wg.Done()
		}()
	}

	wg.Wait()

	expectedCount := 10000
	if count != expectedCount {
		t.Errorf("Expected count to be %d, but got %d", expectedCount, count)
	}
}
