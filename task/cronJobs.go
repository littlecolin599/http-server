package task

import (
	"fmt"
	"log"

	"github.com/RichardKnop/machinery/v1/tasks"
)

func TestPeriodTask() {
	server, _ := startServer()
	signature := &tasks.Signature{
		Name: "periodicTask",
		Args: []tasks.Arg{},
	}
	// 每一分钟执行一次periodicTask函数
	err := server.RegisterPeriodicTask("*/1 * * * ?", "periodic-task", signature)
	if err != nil {
		log.Fatal("Failed to start period task:", err)
		return
	}
	asyncResult, _ := server.SendTask(signature)
	fmt.Println(asyncResult)
}
