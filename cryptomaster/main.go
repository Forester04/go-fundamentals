package main

import (
	"backend/go/crypto/api"
	"fmt"
	"sync"
)

func main() {
	currencies := []string{"BTC", "BCH", "ETH", "USDC", "SATS"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func() {
			GetRateCurrencyData(currency)
			wg.Done()
		}()
	}

	wg.Wait()
}

func GetRateCurrencyData(currency string) {
	rate, err := api.GetRateCurrency(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %d \n", rate.Currency, rate.Price)
	}
}
