package main

import (
	"log"
	"sync"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool {
	return c.logAllowed
}

func NewConfig(allowd bool) config {
	return config{logAllowed: allowd}
}

func main() {

	rps := 1000
	wg := sync.WaitGroup{}

	wg.Add(rps)
	config := NewConfig(true)

	for i := 1; i <= rps; i++ {
		go func(idx int) {
			requestHandler(idx)
			if config.LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}

			wg.Done()

		}(i)
	}
	wg.Wait()
}

func requestHandler(requestidx int) {

}
