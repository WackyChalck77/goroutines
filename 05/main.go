package main

import (
	"fmt"
	"sync/atomic"
)

var counter uint32 //положительный счетчик

func main() {
	for i := 0; i < 50; i++ { //50 горутин
		go func() {
			for c := 0; c < 100; c++ { //выполняет 100 раз
				atomic.AddUint32(&counter, 1) //атомарное прибавление единицы
			}
		}()

	}
	fmt.Println("Counter:", counter)
}
