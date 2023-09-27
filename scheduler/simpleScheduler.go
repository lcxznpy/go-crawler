package scheduler

import "go-crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

// 调度器实现任务进入调度器
func (s *SimpleScheduler) Submit(r engine.Request) {
	//如果这里不开goroutine，那么就会被卡住，效率变低
	go func() {
		s.workerChan <- r
	}()
}

// 绑定chan
func (s *SimpleScheduler) ConfigureWorkChan(c chan engine.Request) {
	s.workerChan = c
}
