package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

var urls = []string{
	"https://en.wikipedia.org/wiki/List_of_largest_companies_in_the_United_States_by_revenue",
	"https://en.wikipedia.org/wiki/List_of_countries_by_GDP_(nominal)",
}

func main() {
	os.Mkdir("output", 0755)

	fmt.Println("Running sequential...")
	start := time.Now()
	for i, url := range urls {
		processURL(url, i+1)
	}
	fmt.Printf("Sequential Execution Time: %.2f seconds\n", time.Since(start).Seconds())

	fmt.Println("\nRunning multithreaded...")
	start = time.Now()
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(u string, idx int) {
			defer wg.Done()
			processURL(u, idx)
		}(url, i+1)
	}
	wg.Wait()
	fmt.Printf("Multithreaded Execution Time: %.2f seconds\n", time.Since(start).Seconds())
}

func processURL(url string, index int) {
	fmt.Printf("Processing URL: %s\n", url)
	h
