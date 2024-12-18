
# Gold Price Scraper

This Go program scrapes gold prices from the website [www.goldtraders.or.th](https://www.goldtraders.or.th) and prints the current prices for different types of gold (gold bars and gold jewelry) along with the current date and time.

## Features
- Fetches live gold prices from the website.
- Displays the following prices:
  - Gold Bar Buy Price (ทองคำแท่งรับซื้อ)
  - Gold Bar Sell Price (ทองคำแท่งขายออก)
  - Gold Jewelry Buy Price (ทองรูปพรรณรับซื้อ)
  - Gold Jewelry Sell Price (ทองรูปพรรณขายออก)
- Prints the date and time when the data was fetched.

## Requirements
- Go 1.18+ (to run the program)
- `go-resty` and `goquery` libraries (for web scraping)

## Installation

1. Install Go from [Go Downloads](https://golang.org/dl/).
2. Install dependencies by running:
   ```bash
   go get github.com/go-resty/resty/v2
   go get github.com/PuerkitoBio/goquery
   ```

## Usage

1. Clone this repository or copy the `main.go` file into your project folder.
2. Run the program:
   ```bash
   go run main.go
   ```
3. The output will display the current gold prices and the date and time of the fetch.

### Sample Output:
```
ราคาประจำวัน:  2024-12-18 14:30:45
ทองคำแท่งรับซื้อ : 42,850.00
ทองคำแท่งขายออก : 42,950.00
ทองรูปพรรณรับซื้อ : 42,084.16
ทองรูปพรรณขายออก : 43,450.00
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
