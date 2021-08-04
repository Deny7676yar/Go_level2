package main

import (
	"fmt"
	"github.com/pkg/profile"
	"runtime"
)


func buff(){
	mutex := make(chan struct{},1)
	counter := 0

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	increase := func(done chan<- struct{}) {
		for{
			mutex <- struct{}{}
			counter++
			if counter == 1000{
				done <- struct {}{}
				return
			}
			<- mutex
		}

	}
	done := make(chan struct{})
	for i := 0; i < 1000; i++{
		go increase(done)
		if i % 100 ==0{// на каждом 100м элементе
			// попросить планировщик прекратить выполнение потока и проверить,
			// нет ли других потоков в состоянии готовности
			runtime.Gosched()

		}
	}

	<- done
	fmt.Println(counter)
}

func main()  {
	buff()
}