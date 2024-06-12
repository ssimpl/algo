package channel

import "sync"

func Merge(cs ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		wg := sync.WaitGroup{}

		for _, c := range cs {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}()
		}

		wg.Wait()

		close(out)
	}()

	return out
}
