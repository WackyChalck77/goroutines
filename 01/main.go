package main

import (
	"fmt"
	"sync"
)

func gorout1(wg *sync.WaitGroup) {
	fmt.Println("Hello from goroutine!")
	defer wg.Done() //уменьшаем счетчик на 1
}
func gorout2(wg *sync.WaitGroup) {
	fmt.Println("Hello from second goroutine!")
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Указываем количество горутин, которые нужно дождаться
	go gorout1(&wg)
	go gorout2(&wg)
	wg.Wait()
}
