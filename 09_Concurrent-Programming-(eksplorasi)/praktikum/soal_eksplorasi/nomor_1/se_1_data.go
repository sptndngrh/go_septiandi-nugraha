// Soal 1
// Buatlah sebuah program yang mengambil data dari sebuah REST API dengan ketentuan sebagai berikut:
// a. Menerapkan penggunaan concurrency (goroutine, channel).
// b. Endpoint yang digunakan: https://fakestoreapi.com/products
// c. Contoh output dari program yang dibuat adalah sebagai berikut:

package main

import (
	"fmt"
	"sync"
	"time"
	"encoding/json"
	"net/http"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

func fetchProducts(url string, wg *sync.WaitGroup, ch chan<- Product) {
	defer wg.Done()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range products {
		ch <- p
	}
}

func main() {
	url := "https://fakestoreapi.com/products"
	numWorkers := 5

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	ch := make(chan Product)

	for i := 0; i < numWorkers; i++ {
		go fetchProducts(url, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Loading please wait...")
	time.Sleep(2 * time.Second)
	fmt.Println("=================================")
	fmt.Println("Products data")

	for p := range ch {
		fmt.Printf("title: %s\nprice: %.2f\ncategory: %s\n===\n", p.Title, p.Price, p.Category)
	}
}