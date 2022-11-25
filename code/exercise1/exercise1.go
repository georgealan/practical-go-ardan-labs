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

	if err != nil {
		log.Fatalf("Error %s", err)
	}
	defer file.Close()

	words, err := mostCommon(file, 10)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	for k, v := range words {
		fmt.Printf("This word: %s, appeared in the text %v times\n", k, v)
	}
}

var wordRegex = regexp.MustCompile(`[a-zA-Z]+`)

func mostCommon(r io.Reader, n int) (map[string]int, error) {
	freqs, err := wordFrequency(r)
	if err != nil {
		return nil, err
	}
	return maxWords(freqs, n)
}

func wordFrequency(r io.Reader) (map[string]int, error) {
	scan := bufio.NewScanner(r)
	freqs := make(map[string]int) // word -> count

	for scan.Scan() {
		words := wordRegex.FindAllString(scan.Text(), -1)
		for _, word := range words {
			freqs[strings.ToLower(word)]++
		}
	}
	if err := scan.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}

func maxWords(freqs map[string]int, n int) (map[string]int, error) {
	if len(freqs) == 0 {
		return nil, fmt.Errorf("Empty map")
	}
	mostFrequents := make(map[string]int)
	counter := 0

	for {
		if counter == n {
			break
		}

		maxNumber, maxWord := 0, ""
		for word, count := range freqs {
			if count > maxNumber {
				maxNumber, maxWord = count, word
			}
		}
		mostFrequents[maxWord] = maxNumber
		delete(freqs, maxWord)
		maxNumber, maxWord = 0, ""
		counter++
	}

	return mostFrequents, nil
}

/*
 TODO Change mostCommon to return the most common n words
  (e.g. func montCommon(r io.Reader, n int) ([]string, error))
  and write the return in each line like this: "This word: (example), appeared in the text 10 times"
*/
