package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	//create a channel to receive a signal
	chanSignal := make(chan os.Signal)
	//set up a channel for receiving Sigterm
	signal.Notify(chanSignal, os.Interrupt)

	// create a context to complete gorutine
	ctx, cancel:= context.WithCancel((context.Background()))

	for i := 0; i < 10; i++ {
		go Do(ctx,  time.Millisecond*time.Duration((i+1)*500), fmt.Sprintf("gorutine %d complete", i))
	}

	//waiting for a signal
	<- chanSignal
	//complete a context
	cancel()
	//waiting for completion gorutine
	time.Sleep(time.Second)
	//complete
	os.Exit(-1)
}

//Do checks if the ctx channel is closed
//if closed, displays the massage
//if not closed, then switch to timeToSleep
func Do(ctx context.Context, timeToSleep time.Duration, message string) {
	for {
		select {
		case <-ctx.Done():
			//reports the end of the goroutine
			fmt.Println(message)
			return
		default:
			time.Sleep(timeToSleep)
			}
	}
}