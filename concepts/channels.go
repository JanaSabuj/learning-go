package main

import (
	"fmt"
	"time"
)

func main() {
	selchannel()
	buf()
	ch := make(chan string)
	go count("morning", ch)

	// for {
	// 	val, ok := <-ch
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(val)
	// }

	for val := range ch {
		fmt.Println(val)
	}

	fmt.Println("Bye")
}

// channels are imp for synchronization between goroutines
func count(thing string, ch chan string) {
	for i := 0; i < 5; i++ {
		ch <- thing // code is blocked here until the receiver is ready
		time.Sleep(time.Millisecond * 500)
	}

	close(ch)
}

// buffered channels - D.N block until buffer is full unline normal channels
// which block as soon as 1 element is pushed to the channel and not received
func buf() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"

	msg := <-ch
	fmt.Println(msg)

	msg = <-ch
	fmt.Println(msg)
}

// select in channels - for async receiving
func selchannel() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			c1 <- "Every 500ms"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			c2 <- "Every two seconds"
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
