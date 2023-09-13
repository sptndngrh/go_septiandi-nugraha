// Soal 3
// Buatlah program yang mencetak bilangan kelipatan 3 dengan menerapkan goroutine dan buffer channel!

package main

import (
	"fmt"
	"time"
)


func printMultipleThree(c chan int) {
	for i := 1; ; i++ {
		if i%3 == 0 {
			c <- i
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan int, 10) 

	go printMultipleThree(c)

	const sleepDurationSeconds = 10
	timer := time.NewTimer(sleepDurationSeconds * time.Second)

	for {
		select {
		case multiple := <- c:
			fmt.Printf("%d merupakan kelipatan dari 3\n", multiple)
		
		case <-timer.C:
			fmt.Println("Program berhenti setelah 10 detik")
			return
		}
	}

}