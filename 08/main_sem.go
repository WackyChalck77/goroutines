package main

import (
	"fmt"
	"time"
)

// функция горутины i, jobs, results, semaphore
func worker(id int, jobs <-chan int, results chan<- int, semaphore chan struct{}) {
	//итерируем по каналу
	for j := range jobs {
		semaphore <- struct{}{} //отдаем в канал пустой слот
		fmt.Printf("Worker %d начал выполнение job %d\n", id, j)
		time.Sleep(time.Second * 2)
		fmt.Printf("Worker %d завершил выполнение job %d\n", id, j)
		// в канал результата вносим номер работы, умноженный на 2
		results <- j * 2
		<-semaphore //освобождаем слот в семафоре
		//забираем одну пустую структуру из семафора
	}
}

func main() {
	const maxWorkers = 2
	const numberOfJobs = 12

	jobs := make(chan int, numberOfJobs)
	results := make(chan int, numberOfJobs)
	//создаем семафор, буферизированный канал со структурами
	//максимальный буфер - maxWorkers
	semaphore := make(chan struct{}, maxWorkers)

	for i := 1; i <= maxWorkers; i++ {
		//запускаем горутнины, они принимают
		//номер от 1 до максимального количества одновременно
		//работающих горутин, каналы:с работой, результатом, семафором
		go worker(i, jobs, results, semaphore)
	}
	//в канал с работой прилетает номер работы
	for j := 1; j <= numberOfJobs; j++ {
		jobs <- j //значение счетчика отправляем в канал jobs
	}
	close(jobs) //закрываем канал с работой

	for k := 1; k <= numberOfJobs; k++ {
		<-results //просто ждём сигнала о завершении из канала
		//results
	}
}
