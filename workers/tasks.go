package workers

import (
	"fmt"
	"time"
)

func Add(args ...int64) (int64, error) {
	println("############# 执行耗时任务 #############")
	time.Sleep(10 * time.Second)
	sum := int64(0)
	for _, arg := range args {
		sum += arg
	}
	println("############# 耗时任务Done #############")
	return sum, nil
}

func PeriodicTask() error {
	fmt.Println("################ 执行周期任务 #################")
	return nil
}
