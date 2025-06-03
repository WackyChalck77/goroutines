package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// приходит задача - > бросаем ее в канал -> воркеры сами разбираются
// количество одновременно работающих горутин
const NumberOfGoroutines = 2

// горутина worker принимает порядковый номер, канал с задачей для чтения
// задачей будет возведение числа в степень
func Worker(wId int, taskChan <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range taskChan {
		time.Sleep(time.Millisecond * 300)
		fmt.Println(strings.Repeat("---", wId), "Worker", wId,
			"отработал задачу", task, "результат", task*task)
	}
}

func main() {
	wg := sync.WaitGroup{}
	taskChan := make(chan int, 3) //можем менять буфер
	//запускаем воркеров. Количество берем из константы
	for i := 1; i <= NumberOfGoroutines; i++ {
		wg.Add(1)
		go Worker(i, taskChan, &wg)
	}
	taskData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//заполняем канал с задачами самими задачами (слайсом)
	for _, task := range taskData {
		taskChan <- task
	}
	close(taskChan)
	wg.Wait()

	//	time.Sleep(time.Second * 2)
}
