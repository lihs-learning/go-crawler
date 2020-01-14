package scheduler

import "github.com/lihs-learning/go-crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(in chan engine.Request) {
	s.workerChan = in
}
