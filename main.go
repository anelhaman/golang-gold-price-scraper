package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type GoldPrices struct {
	BaseURL string
}

func (gp *GoldPrices) FetchPrices() {
	client := resty.New()

	// Fetch HTML content
	resp, err := client.R().Get(gp.BaseURL)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	// Parse the HTML response
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		log.Fatalf("Failed to parse HTML: %v", err)
	}

	// A map to store extracted prices
	prices := sync.Map{}

	// WaitGroup to handle concurrent extraction
	var wg sync.WaitGroup

	// Function to extract a price and store it in the map
	extractPrice := func(key, selector string) {
		defer wg.Done() // Signal WaitGroup when done

		// Find the price
		price := doc.Find(selector).Text()

		// Store in sync.Map
		prices.Store(key, price)
	}

	// Add tasks to WaitGroup and start goroutines
	wg.Add(4)
	go extractPrice("ทองคำแท่งรับซื้อ", "#DetailPlace_uc_goldprices1_lblBLBuy b font")
	go extractPrice("ทองคำแท่งขายออก", "#DetailPlace_uc_goldprices1_lblBLSell b font")
	go extractPrice("ทองรูปพรรณรับซื้อ", "#DetailPlace_uc_goldprices1_lblOMBuy b font")
	go extractPrice("ทองรูปพรรณขายออก", "#DetailPlace_uc_goldprices1_lblOMSell b font")

	// Wait for all goroutines to finish
	wg.Wait()

	// Print current date and time
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	fmt.Println("ราคาประจำวัน: ", currentTime)

	// Print results in simple format
	prices.Range(func(key, value interface{}) bool {
		fmt.Printf("%s : %s\n", key, value)
		return true
	})
}

func main() {
	gp := &GoldPrices{BaseURL: "https://www.goldtraders.or.th"}
	gp.FetchPrices()
}
