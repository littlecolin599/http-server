package task

import "log"

func Start() {
	// 启动worker
	if err := worker(); err != nil {
		log.Fatal("Failed to start, err:", err)
	}
}

func StartCron() {
	TestPeriodTask()
}
