package timers

import (
	"fmt"
	"time"
)

func timers(){
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	stop2 := timer2.Stop() // This would never provide chance for timer 2 to run

	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}

func tickers(){
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
				case <-done:
					return
				case t := <-ticker.C:
					fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker Stopped")
}

func TimersExamples() {
	timers()
	tickers()
}