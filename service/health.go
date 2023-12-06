package service

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/emersonfbarros/backend-challenge-klever/config"
	"github.com/emersonfbarros/backend-challenge-klever/model"
)

type ExtApi struct {
	Status       string `json:"status"`
	ResponseTime string `json:"responseTime"`
}

type HealthRes struct {
	Status      string `json:"status"`
	Timestamp   string `json:"timestamp"`
	Uptime      string `json:"uptime"`
	ExternalApi ExtApi `json:"externalApi"`
}

func (s *Services) Health(fetcher model.Fetcher) *HealthRes {
	addressTest := os.Getenv("ADDRESS_TEST")
	txTest := os.Getenv("TX_TEST")
	successCount := 0
	var wg sync.WaitGroup
	wg.Add(3)

	startTime := time.Now()

	go func() {
		defer wg.Done()
		_, err := fetcher.Fetch("address", addressTest)
		if err == nil {
			successCount++
		}
	}()

	go func() {
		defer wg.Done()
		_, err := fetcher.Fetch("utxo", addressTest)
		if err == nil {
			successCount++
		}
	}()

	go func() {
		defer wg.Done()
		_, err := fetcher.Fetch("tx", txTest)
		if err == nil {
			successCount++
		}
	}()

	wg.Wait()

	extApiResTime := strconv.FormatInt(time.Since(startTime).Milliseconds(), 10)

	var extApiStatus string
	if successCount == 3 {
		extApiStatus = "Ok"
	}
	if successCount == 2 || successCount == 1 {
		extApiStatus = "Partially Ok"
	}
	if successCount == 0 {
		extApiStatus = "Down"
	}

	uptimeDuration := time.Since(config.AppStartTime)

	health := HealthRes{
		Status:    "healthy",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime: fmt.Sprintf("%dD %dH %dM %dS",
			int(uptimeDuration.Hours()/24),
			int(uptimeDuration.Hours())%24,
			int(uptimeDuration.Minutes())%60,
			int(uptimeDuration.Seconds())%60),
		ExternalApi: ExtApi{
			Status:       extApiStatus,
			ResponseTime: fmt.Sprintf("%sms", extApiResTime),
		},
	}

	return &health
}
