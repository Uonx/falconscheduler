package scheduler

import (
	"github.com/Uonx/falconscheduler/pkg/engine"
)

type QueuedScheduler struct {
	itemChan   chan engine.WorkerMethod
	workerChan chan chan engine.WorkerMethod
}

func (s *QueuedScheduler) WorkerChan() chan engine.WorkerMethod {
	return make(chan engine.WorkerMethod)
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.WorkerMethod)
	s.itemChan = make(chan engine.WorkerMethod)
	go func() {
		var itemQ []engine.WorkerMethod
		var workerQ []chan engine.WorkerMethod
		for {
			var activeItem engine.WorkerMethod
			var activeWorker chan engine.WorkerMethod
			if len(itemQ) > 0 && len(workerQ) > 0 {
				activeItem = itemQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case i := <-s.itemChan:
				itemQ = append(itemQ, i)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
			case activeWorker <- activeItem:
				itemQ = itemQ[1:]
				workerQ = workerQ[1:]
			default:
			}
		}
	}()
}

func (s *QueuedScheduler) WorkerReady(w chan engine.WorkerMethod) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Submit(i engine.WorkerMethod) {
	s.itemChan <- i
}
