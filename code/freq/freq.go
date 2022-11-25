package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	file, err := os.Open("freq/sherlock.txt")
	// In Golang IDE use this path: freq/sherlock.txt
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	defer file.Close()

	w, err := mostCommon(file)
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
	fmt.Println("The most common word is:", w)

	//mapDemo()
}

// "Who's on first?" -> [Who s on first]
var wordReg = regexp.MustCompile(`[a-zA-Z]+`) // This will run before main

func mostCommon(r io.Reader) (string, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return "", err
	}

	return maxWord(freqs)
}

func mapDemo() {
	// Always is a good idea comment your maps to clarify you intent.
	var stocks map[string]float64 // symbol -> price
	symbol := "TTWO"
	price := stocks[symbol]
	fmt.Printf("%s -> $%.2f\n", symbol, price)

	if price, ok := stocks[symbol]; ok {
		fmt.Printf("%s -> $%.2f\n", symbol, price)
	} else {
		fmt.Printf("%s not found\n", symbol)
	}

	stocks = map[string]float64{
		symbol: 137.73,
		"AAPL": 172.35,
	}

	for k := range stocks { // keys
		fmt.Println("Only keys:", k)
	}

	for k, v := range stocks { // key & value
		fmt.Println(k, "->", v)
	}

	for _, v := range stocks { // only values
		fmt.Println("only values:", v)
	}

	delete(stocks, "AAPL")
	fmt.Println(stocks)
	delete(stocks, "AAPL") // Not panic
}

/* You can use this way too, to run code before main
func init()  {
	// ...
}
*/

/* Multi lines strings You can use raw strings to create multi line strings
var request = `GET /ip HTTP/1.1
Host: httpbin.org
Connection: Close
`
*/

func maxWord(freqs map[string]int) (string, error) {
	if len(freqs) == 0 {
		return "", fmt.Errorf("Empty map")
	}

	maxN, maxW := 0, ""
	for word, count := range freqs {
		if count > maxN {
			maxN, maxW = count, word
		}
	}
	return maxW, nil
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count

	for s.Scan() {
		words := wordReg.FindAllString(s.Text(), -1) // current line
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}
