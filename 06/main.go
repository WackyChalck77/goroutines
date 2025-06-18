package main

import (
	"context"
	"fmt"
	"time"
)

// func worker(ctx context.Context, workerNum int, out chan<- int) {
// 	waitTime := time.Duration(rand.Intn(100)) * time.Millisecond
// 	fmt.Println(workerNum, "спит на", waitTime, "milliseconds")
// 	select {
// 	case <-ctx.Done():
// 		return
// 	case <-time.After(waitTime):
// 		fmt.Println("worker", workerNum, "завершен")
// 		out <- workerNum
// 	}
// }

func StartBatchProcessor(ctx context.Context, input <-chan int) {
	const maxBatch = 5
	const n = 2
	const waitTime = n * time.Second
	//waitTime := time.Duration(rand.Intn(100)) * time.Millisecond
	batch := make([]int, 0, maxBatch) //заготавливаем слайс len, capacity
	timer := time.NewTimer(waitTime)  //ставим таймер

	//If you omit the loop condition it loops
	// forever, so an infinite loop is compactly expressed.
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Контекст завершен извне")
			return
		case inputGoes := <-input:
			batch = append(batch, inputGoes)
			fmt.Println("Текущий бэтч:", batch)
			if len(batch) == maxBatch {
				fmt.Println("Заполнен бэтч", batch)
				batch = batch[:0]  //это обнуление бэтча
				if !timer.Stop() { //если таймер не на стопе
					<-timer.C //блокирует выполнение пока не придет событие из канала
					//timer, C - канал, который используется для отправки уведомления
					//о срабатывании таймера.
				}
				timer.Reset(waitTime)
			}
		case <-timer.C:
			fmt.Println("Текущий бэтч, таймер завершен", batch)
			batch = batch[:0] //это обнуление бэтча
			timer.Reset(waitTime)

		}
	}
}

func main() {
	// инициализация канала
	/* создание контекста  */
	ctx, finish := context.WithCancel(context.Background())
	chan1 := make(chan int, 5)

	go func() {
		for i := 1; i < 21; i++ {
			chan1 <- i
			time.Sleep(10 * time.Millisecond)
		}
		//finish()
	}()
	go StartBatchProcessor(ctx, chan1)

	// for i := 0; i < 6; i++ {
	// 	go worker(ctx, i, chan1)
	// }
	// found := <-chan1
	// fmt.Println("Результат пришел, номер воркера", found)
	time.Sleep(2500 * time.Millisecond)

	finish() //остановка работы всех горутин

	// сбор данных

	fmt.Println("Main: процессы завершены")
}
