package scheduler

import "go-crawler/engine"

type QueueScheduler struct {
	//这两个chan是全局的
	RequestChan chan engine.Request      //记录任务通道
	WorkerChan  chan chan engine.Request // 记录工人空闲通道的通道，每个工人都有一个单独的通道接收任务
}

// 初始化请求队列
func (q *QueueScheduler) WorkChan() chan engine.Request {
	return make(chan engine.Request)
}

// 提交任务
func (q *QueueScheduler) Submit(r engine.Request) {
	q.RequestChan <- r
}

// 工人的chan 提交
func (q *QueueScheduler) WorkReady(w chan engine.Request) {
	q.WorkerChan <- w
}

func (q *QueueScheduler) Run() {
	//创建空闲工人队列chan
	q.WorkerChan = make(chan chan engine.Request) //通道的使用一定要初始化
	//创建任务队列chan
	q.RequestChan = make(chan engine.Request) //通道的使用一定要初始化
	go func() {
		//任务队列，接收到任务就会放进去
		var requestQueue []engine.Request
		//空闲工人队列，接收到工人就会放进去
		var workQueue []chan engine.Request
		for {
			//第一个任务
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			//第一个空闲的工人
			//如果都有元素，就分配第一个任务给第一个工人
			if len(requestQueue) > 0 && len(workQueue) > 0 {
				activeWorker = workQueue[0]
				activeRequest = requestQueue[0]
			}
			select {
			//有请求，就放入请求队列
			case r := <-q.RequestChan:
				requestQueue = append(requestQueue, r)
				//有空闲工人chan，就放入空闲工人chan队列
			case w := <-q.WorkerChan:
				workQueue = append(workQueue, w)
				//如果有任务和有空闲工人，则将任务分配给第一个空闲工人in队列  处理
			case activeWorker <- activeRequest:
				//删除两个队列的第一个元素
				workQueue = workQueue[1:]
				requestQueue = requestQueue[1:]
			}
		}
	}()
}
