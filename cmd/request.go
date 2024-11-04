package cmd

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

func performRequest(url string, wg *sync.WaitGroup, results chan<- *http.Response) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Erro ao fazer request para %s: %s\n", url, err)
		return
	}

	results <- resp
}

func executeLoadTest(url string, totalRequests int, concurrency int, timeout time.Duration) {
	var wg sync.WaitGroup
	results := make(chan *http.Response, totalRequests)
	sem := make(chan struct{}, concurrency)
	requestCount := int32(0)

	startTime := time.Now()

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func() {
			defer wg.Done()
			defer func() { <-sem }()

			client := &http.Client{
				Timeout: timeout,
			}
			resp, err := client.Get(url)
			if err != nil {
				fmt.Printf("Erro ao fazer request para %s: %s\n", url, err)
				return
			}

			atomic.AddInt32(&requestCount, 1)
			results <- resp
		}()
	}

	wg.Wait()
	close(results)

	totalDuration := time.Since(startTime)
	fmt.Printf("Tempo total de execução: %s\n", totalDuration)
	fmt.Printf("Total de requests enviados: %d\n", requestCount)

	processResults(results)
}

func processResults(results chan *http.Response) {
	statusCount := make(map[int]int)

	for resp := range results {
		statusCount[resp.StatusCode]++
		resp.Body.Close()
	}

	for status, count := range statusCount {
		fmt.Printf("Status %d: %d\n", status, count)
	}
}
