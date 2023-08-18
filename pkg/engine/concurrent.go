package engine

type ReadyNotifier interface {
	WorkerReady(chan WorkerMethod)
}

type Scheduler interface {
	ReadyNotifier
	Submit(WorkerMethod)
	WorkerChan() chan WorkerMethod
	Run()
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

func (e *ConcurrentEngine) Run(seeds ...WorkerMethod) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func(item Item) { e.ItemChan <- item }(item)
			// e.ItemChan <- item
		}

		for _, work := range result.Workers {
			e.Scheduler.Submit(work)
		}
	}
}

func (e *ConcurrentEngine) createWorker(in chan WorkerMethod, out chan ParseResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in)

			request := <-in
			result, err := request.Parse()
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
