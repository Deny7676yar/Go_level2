package main

import (
	"fmt"
	"lesson4/workerpool"
	"time"
)

func main() {
	var allTask []*workerpool.Task
	val := 0
	for i := 1; i <= 1000; i++ {
		task := workerpool.NewTask(func(data interface{}) error {
			//taskID := data.(int)
			time.Sleep(100 * time.Millisecond)
			//fmt.Printf("Task %d processed\n", taskID)
			val +=1
			return nil
		}, i)
		allTask = append(allTask, task)
	}

	pool := workerpool.NewPool(allTask, 5)
	pool.Run()


	fmt.Printf("Value = %d\n", val)
}