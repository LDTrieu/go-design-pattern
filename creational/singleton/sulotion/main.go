package main

import (
	"log"
	"sync"
	"time"
)

type config struct {
	logAllowed bool
}

func (c config) LogAllowed() bool { return c.logAllowed }

func NewConfig(allowed bool) config {
	return config{logAllowed: allowed}
}

var singletonApp = &application{once: sync.Once{}}

func GetApplication() *application {
	return singletonApp
}

type application struct {
	once sync.Once
	cfg  *config
}

func (app *application) GetConfig() *config {

	if app.cfg == nil {
		app.once.Do(func() {
			log.Println("Loading config once and forever.")
			app.loadConfig()
		})
	}

	return app.cfg
}

func (app *application) loadConfig() {
	time.Sleep(100)
	app.cfg = &config{logAllowed: true}
}

func main() {
	rps := 1000
	wg := sync.WaitGroup{}
	wg.Add(rps)

	for i := 1; i <= rps; i++ {
		go func(idx int) {
			requestHandler(idx)

			if GetApplication().GetConfig().LogAllowed() {
				log.Printf("Request %d handled successfully.\n", idx)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
}

func requestHandler(requestIdx int) {

	if GetApplication().GetConfig().LogAllowed() {
		log.Printf("Handling request %d... please wait.\n", requestIdx)
	}
}
