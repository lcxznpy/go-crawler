package scheduler

import "go-crawler/engine"

type QueueScheduler struct {
	RequestChan chan engine.Request
	WorkerChan  chan chan engine.Request // 记录工人空闲通道的通道
}

func (q *QueueScheduler) Submit(r engine.Request) {
	q.RequestChan <- r
}

func (q *QueueScheduler) ConfigureWorkChan(r chan engine.Request) {
	//TODO implement me
	panic("implement me")
}

func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.WorkerChan <- w
}

func (q *QueueScheduler) Run() {
	var requestQueue []engine.Request
	var workQueue []chan engine.Request
	go func() {
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workQueue) > 0 {
				activeWorker = workQueue[0]
				activeRequest = requestQueue[0]
			}
			select {
			case r := <-q.RequestChan:
				requestQueue = append(requestQueue, r)
			case w := <-q.WorkerChan:
				workQueue = append(workQueue, w)
			case activeWorker <- activeRequest:
				workQueue = workQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
