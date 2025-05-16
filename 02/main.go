package main

import (
	"fmt"
	"sync"
)

func gorout1(wg *sync.WaitGroup) {
	fmt.Println("goroutine 1")
	defer wg.Done() //уменьшаем счетчик на 1
}
func gorout2(wg *sync.WaitGroup) {
	fmt.Println("goroutine 2")
	defer wg.Done() //уменьшаем счетчик на 1
}
func gorout3(wg *sync.WaitGroup) {
	fmt.Println("goroutine 3")
	defer wg.Done() //уменьшаем счетчик на 1
}
func gorout4(wg *sync.WaitGroup) {
	fmt.Println("goroutine 4")
	defer wg.Done() //уменьшаем счетчик на 1
}
func gorout5(wg *sync.WaitGroup) {
	fmt.Println("goroutine 5")
	defer wg.Done() //уменьшаем счетчик на 1
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go gorout1(&wg)
	go gorout2(&wg)
	go gorout3(&wg)
	go gorout4(&wg)
	go gorout5(&wg)
	wg.Wait()
}
