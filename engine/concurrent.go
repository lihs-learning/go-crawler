package engine

import "log"

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	Persist Persist
}

type Scheduler interface {
	ReadyNotifier

	Submit(request Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	invalidSeedNum := 0
	for _, seed := range seeds {
		if isDuplicate(seed.URL) {
			invalidSeedNum++
			continue
		}
		e.Scheduler.Submit(seed)
	}
	log.Printf("invlide seed num: %d", invalidSeedNum)

	for {
		result := <-out

		for _, item := range result.Items {
			e.Persist.Save(item)
		}

		for _, request := range result.Requests {
			if isDuplicate(request.URL) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, scheduler Scheduler) {
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
