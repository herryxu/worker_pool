package main

import (
	"fmt"
	"math/rand"
	"time"
	"workerpool/workerpool"
)

var poolClient = &workerpool.Pool{}

func main() {
	poolClient = workerpool.NewPool(5)
	go func() {
		i := 0
		for {
			taskID := i

			if i == 10 {
				break
			}
			task := workerpool.NewTask(func(data interface{}) error {
				taskID := data.(int)
				rSed := rand.Intn(5)
				time.Sleep(time.Duration(rSed) * time.Second)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			poolClient.AddTask(task)
			i++
		}
	}()
	go func() {
		newTask()
	}()
	poolClient.RunBackground()
	time.Sleep(20 * time.Second)
	poolClient.Stop()
}
func newTask() {
	i := 10
	for {
		taskID := i

		if i == 20 {
			break
		}
		task := workerpool.NewTask(func(data interface{}) error {
			taskID := data.(int)
			rSed := rand.Intn(5)
			time.Sleep(time.Duration(rSed) * time.Second)
			fmt.Printf("Task %d processed\n", taskID)
			return nil
		}, taskID)
		poolClient.AddTask(task)
		i++
	}
	fmt.Printf("%+v", poolClient)
}
