package scheduler

import (
	"polaris/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workChan    chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}
func (q *QueuedScheduler) WorkerReady(w chan engine.Request) {
	q.workChan <- w
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) Run() {
	q.workChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)
	go func() {
		for {
			var requestQueue [] engine.Request
			var workerQueue [] chan engine.Request

			var activeRequest engine.Request
			var activeWorker chan engine.Request

			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case r := <-q.requestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.workChan:
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				workerQueue = workerQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
