package main

import (
	"sync"
)

type httpPkg struct{}

func (httpPkg) Get(url string) {}

var http httpPkg

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// увеличиваем счетчик WaitGroup
		wg.Add(1)
		// стартуем горутину для выборки по URL
		go func(url string) {
			// уменьшаем счетчик при завершении горутины
			defer wg.Done()
			// выбираем данные по URL
			http.Get(url)
		}(url)
	}
	// ожидаем, пока все запросы не будут завершены
	wg.Wait()

}
