package task

import (
	"http-server/workers"
	"log"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
)

var AsyncTaskCenter *machinery.Server

func startServer() (*machinery.Server, error) {
	// 配置Machinery
	serverCfg, err := config.NewFromYaml("./config/machinery_config.yml", true)
	if err != nil {
		log.Fatal("Failed to read configuration:", err)
	}
	server, err := machinery.NewServer(serverCfg)
	if err != nil {
		log.Fatal("Failed to create server:", err)
	}

	taskMap := initAsyncTaskMap()
	AsyncTaskCenter = server
	return server, server.RegisterTasks(taskMap)

}

// 启动worker
func worker() error {
	consumerTag := "machine_worker"
	server, err := startServer()
	if err != nil {
		log.Fatal("Failed to start server:", err)
		return nil
	}
	worker := server.NewWorker(consumerTag, 0) // 第二个参数表示并发数，0表示不限制

	// 钩子函数
	errorHandler := func(err error) {}
	pretaskhandler := func(signature *tasks.Signature) {}
	posttaskhandler := func(signature *tasks.Signature) {}

	worker.SetPostTaskHandler(posttaskhandler)
	worker.SetPreTaskHandler(pretaskhandler)
	worker.SetErrorHandler(errorHandler)

	return worker.Launch()
}

// 注册任务
func initAsyncTaskMap() map[string]interface{} {
	taskMap := map[string]interface{}{
		"add":          workers.Add,
		"periodicTask": workers.PeriodicTask,
	}
	return taskMap
}
