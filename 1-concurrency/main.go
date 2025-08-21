package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	ch2 := doubler(generator())
	for v := range ch2 {
		fmt.Println(v)
	}

}

func generator() <-chan int {
	ch := make(chan int)

	go func() {
		for range 10 {
			ch <- rand.IntN(100)
		}

		close(ch)
	}()

	return ch
}

func doubler(inputCh <-chan int) <-chan int {
	ch2 := make(chan int)

	go func() {
		for v := range inputCh {
			ch2 <- v * v
		}

		close(ch2)
	}()

	return ch2
}
