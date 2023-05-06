package task

import (
	"fmt"

	"github.com/RichardKnop/machinery/v1/tasks"
)

/* 触发执行Add异步任务 */
func TaskAdd(a, b int64) {
	signature := &tasks.Signature{
		Name: "add",
		Args: []tasks.Arg{
			{
				Type:  "int64",
				Value: a,
			},
			{
				Type:  "int64",
				Value: b,
			},
		},
	}
	_, err := AsyncTaskCenter.SendTask(signature) // 任务可以通过将Signature的实例传递给Server实例来调用
	if err != nil {
		fmt.Println(err)
	}
}
