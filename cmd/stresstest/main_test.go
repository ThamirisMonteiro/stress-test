package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type RequestData struct {
	StatusCode   int
	ResponseTime time.Duration
}

func ExecuteRequests(url string, count int) ([]RequestData, error) {
	var data []RequestData
	for i := 0; i < count; i++ {
		start := time.Now()
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		responseTime := time.Since(start)
		data = append(data, RequestData{
			StatusCode:   resp.StatusCode,
			ResponseTime: responseTime,
		})
		resp.Body.Close()
	}
	return data, nil
}

func GenerateReport(data []RequestData) struct {
	TotalRequests int
	SuccessCount  int
	AverageTime   time.Duration
} {
	totalRequests := len(data)
	successCount := 0
	totalTime := time.Duration(0)

	for _, entry := range data {
		if entry.StatusCode == http.StatusOK {
			successCount++
		}
		totalTime += entry.ResponseTime
	}

	averageTime := totalTime / time.Duration(totalRequests)

	return struct {
		TotalRequests int
		SuccessCount  int
		AverageTime   time.Duration
	}{
		TotalRequests: totalRequests,
		SuccessCount:  successCount,
		AverageTime:   averageTime,
	}
}

func mockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRequestExecution(t *testing.T) {
	server := mockServer()
	defer server.Close()

	data, err := ExecuteRequests(server.URL, 10)
	if err != nil {
		t.Fatalf("Erro ao executar requests: %v", err)
	}

	if len(data) != 10 {
		t.Errorf("Esperava 10 dados, mas recebi %d", len(data))
	}
}

func TestDataCollection(t *testing.T) {
	server := mockServer()
	defer server.Close()

	data, err := ExecuteRequests(server.URL, 100)
	if err != nil {
		t.Fatalf("Erro ao executar requests: %v", err)
	}

	if len(data) != 100 {
		t.Errorf("Esperava 100 dados, mas recebi %d", len(data))
	}

	for _, entry := range data {
		if entry.StatusCode != http.StatusOK {
			t.Errorf("Status inesperado: %d", entry.StatusCode)
		}
	}
}

func TestReportGeneration(t *testing.T) {
	server := mockServer()
	defer server.Close()

	data, err := ExecuteRequests(server.URL, 100)
	if err != nil {
		t.Fatalf("Erro ao executar requests: %v", err)
	}
	report := GenerateReport(data)

	if report.TotalRequests != 100 {
		t.Errorf("Esperava 100 requisições no relatório, mas recebi %d", report.TotalRequests)
	}
	if report.SuccessCount != 100 {
		t.Errorf("Esperava 100 requisições bem-sucedidas, mas recebi %d", report.SuccessCount)
	}
}

func TestAverageResponseTime(t *testing.T) {
	server := mockServer()
	defer server.Close()

	data, err := ExecuteRequests(server.URL, 10)
	if err != nil {
		t.Fatalf("Erro ao executar requests: %v", err)
	}

	report := GenerateReport(data)

	if report.AverageTime == 0 {
		t.Error("Esperava um tempo médio de resposta maior que zero.")
	}
}

func TestHandleErrorResponses(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	data, err := ExecuteRequests(server.URL, 5)
	if err != nil {
		t.Fatalf("Erro ao executar requests: %v", err)
	}

	report := GenerateReport(data)

	if report.TotalRequests != 5 {
		t.Errorf("Esperava 5 requisições, mas recebi %d", report.TotalRequests)
	}
	if report.SuccessCount != 0 {
		t.Errorf("Esperava 0 requisições bem-sucedidas, mas recebi %d", report.SuccessCount)
	}
}

func TestInvalidURL(t *testing.T) {
	data, err := ExecuteRequests("http://invalid.url", 5)
	if err == nil {
		t.Fatal("Esperava um erro ao acessar uma URL inválida.")
	}

	if data != nil {
		t.Errorf("Esperava nil para dados, mas recebi %v", data)
	}
}
