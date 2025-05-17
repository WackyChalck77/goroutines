package main

import "fmt"

func oneFive(ch chan int) {
	for i := 1; i < 6; i++ {
		ch <- i

	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go oneFive(ch)
	sum := 0
	for v := range ch {
		sum = sum + v
		fmt.Println(v, sum)
	}
}
