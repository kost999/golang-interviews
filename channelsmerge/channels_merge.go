package channelsmerge

import (
	"sync"
)

func Merge(chs ...chan int) chan int {
	r := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(len(chs))
	for _, ch := range chs {
		go func() {
			defer wg.Done()
			for v := range ch {
				r <- v
			}
		}()
	}

	go func() {
		defer close(r)
		wg.Wait()
	}()

	return r
}
