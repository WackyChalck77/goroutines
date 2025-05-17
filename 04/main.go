package main

import (
	"fmt"
	"sync"
)

type Counter struct { //структура
	mu sync.Mutex
	v  int
}

func (c *Counter) Inc() { //метод
	c.mu.Lock()
	c.v++
	fmt.Println(c.v)
	c.mu.Unlock()
}

func main() {
	//	var wg sync.WaitGroup
	c := &Counter{v: 0}

	for i := 0; i < 10; i++ {
		//	wg.Add(1)
		go func() {

			c.Inc()
			//		defer wg.Done()
		}()
	}

	//wg.Wait()

	fmt.Println("Done %d", c.v)
}
