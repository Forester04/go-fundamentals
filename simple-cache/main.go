package main

import (
	"fmt"
	"simplecache/woo/cache/simplecache"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	store := &simplecache.ConcurrentStore{
		Store: make(map[string]any),
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Set("firstname", "Erick")
		fmt.Println("First Goroutine: Added successfully")
		fmt.Printf("Current store: %+v\n", store.Store)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Set("lastname", "Forester")
		store.Set("type", 1)
		fmt.Println("Second Goroutine: Added successfully")
		fmt.Printf("Current store: %+v\n", store.Store)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		value, _ := store.Get("firstname")
		fmt.Printf("The first value read is %v\n", value)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		value, _ := store.Get("type")
		fmt.Printf("The second value read is %v\n", value)
	}()

	wg.Wait()
	fmt.Println("✅ All goroutines completed successfully.")
}
