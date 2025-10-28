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

	wg.Go(func() {
		store.Set("firstname", "Erick")
		fmt.Printf("First item added successfully: %v\n", store.Store)
	})

	wg.Go(func() {
		store.Set("lastname", "Forester")
		fmt.Printf("Second item added successfully: %v\n", store.Store)
	})

	wg.Go(func() {
		value, _ := store.Get("firstname")
		fmt.Printf("First retrieved value: %v\n", value)
	})
	wg.Go(func() {
		value, _ := store.Get("lastname")
		fmt.Printf("First retrieved value: %v\n", value)
	})

	wg.Wait()
	fmt.Println("✅ Program exited successfully")
}
