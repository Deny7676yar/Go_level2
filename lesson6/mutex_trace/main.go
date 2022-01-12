package main

import (
	"fmt"
	"github.com/pkg/profile"
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
	}

	<- done
	fmt.Println(counter)
}

func main()  {
	buff()
}