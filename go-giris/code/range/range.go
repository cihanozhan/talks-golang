package main

import (
	"fmt"
	"time"
)

// placeholder so we have the time package imported, so it doesn't break
// when we uncomment the time.Sleep statement inside the fibonacci
// function
var b = time.Second

// START OMIT
func fibonacci(c chan int) {
	x, y := 0, 1
	for i := 0; i < 7; i++ {
		c <- x
		// slow producer
		// time.Sleep(time.Millisecond * 300)
		x, y = y, x+y
	}
	fmt.Println("finished")
	close(c)
}

func main() {
	c := make(chan int)

	go fibonacci(c) // collect numbers

	for i := range c {
		// slow consumer
		// time.Sleep(time.Millisecond * 300)
		fmt.Println(i)
	}
}

// END OMIT
