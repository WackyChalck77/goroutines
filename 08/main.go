package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// функция принимает слайс адресов,
// выдает мапу url : содержимое ответа
func FetchURLs(urls []string) map[string]string {
	//объявляем средства работы с многопоточностью
	wg := sync.WaitGroup{}
	var mu sync.Mutex

	// client := &http.Client{ //клиент с таймаутом
	// 	Timeout: 10 * time.Second,
	// }
	//отмена контекстом по истечении таймера
	n := 2200 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), n)
	defer cancel()
	result := make(map[string]string)
	for _, url := range urls {
		//fmt.Println(url)
		//создадим горутину для обработки http-запроса для каждого адреса
		wg.Add(1)
		go func(url string) {
			defer wg.Done() //по окончании уменьшаем счетчик

			select {
			case <-ctx.Done():
				fmt.Println("контекст завершился быстрее горутины")
				return
			default:
				//станартный запрос
				//	resp, err := http.Get(url)
				//используем Client с таймаутом
				// resp, err := client.Get(url)
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					fmt.Println("Ошибка при переборе url")
					mu.Lock()
					result[url] = "error"
					mu.Unlock()
					return
				}
				req = req.WithContext(ctx)

				resp, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Println("Ошибка при выполнении запроса")
					mu.Lock()
					result[url] = "error"
					mu.Unlock()
					return
				}
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("Ошибка при чтении body в url")
					mu.Lock()
					result[url] = "error"
					mu.Unlock()
					return
				}

				bodyString := string(body[:34])
				//fmt.Println(bodyString)
				mu.Lock()
				result[url] = bodyString
				mu.Unlock()
				resp.Body.Close()
				// Собирает результаты (код ответа и часть тела) в map[string]string,
				// где: ключ — URL
				// значение — содержимое ответа (ограниченное, например, 100 символами)
			}
		}(url)
	}

	wg.Wait()
	return result
}

func main() {
	urls := []string{"https://yandex.ru", "https://google.com", "https://bing.com"}

	result := FetchURLs(urls)
	for url, content := range result {
		fmt.Printf("URL: %s \t CONTENT: %s\n", url, content)
	}
}
