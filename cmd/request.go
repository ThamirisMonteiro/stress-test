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
	status200Count := int32(0)
	statusCounts := make(map[int]int)

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
			if resp.StatusCode == http.StatusOK {
				atomic.AddInt32(&status200Count, 1)
			} else {
				statusCounts[resp.StatusCode]++
			}
			results <- resp
		}()
	}

	wg.Wait()
	close(results)

	totalDuration := time.Since(startTime)

	generateReport(totalDuration, requestCount, status200Count, statusCounts)
}

func generateReport(totalDuration time.Duration, totalRequests int32, status200Count int32, statusCounts map[int]int) {
	fmt.Println("\n--- Relatório de Teste de Carga ---")
	fmt.Println("-----------------------------------")
	fmt.Printf("Tempo total de execução: %.2fs\n", totalDuration.Seconds())
	fmt.Printf("Total de requests enviados: %d\n", totalRequests)
	fmt.Printf("Requests com status 200: %d\n", status200Count)

	fmt.Println("\nDistribuição de outros códigos de status:")
	for code, count := range statusCounts {
		fmt.Printf("Código %d: %d\n", code, count)
	}
	fmt.Println("-----------------------------------")
	fmt.Println("Teste concluído com sucesso.")
	fmt.Println("-----------------------------------")
}
