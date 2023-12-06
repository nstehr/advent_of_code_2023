package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	input = "input.txt"
)

func main() {
	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		text := scanner.Text()
		first := ""
		firstIdx := math.MaxInt
		last := ""
		lastIdx := 0
		log.Println(text)
		// more elegant ways to solve this, but this works
		for k, v := range numMap {
			idx := strings.Index(text, k)
			endIdx := strings.LastIndex(text, k)

			if idx > -1 {
				if idx <= firstIdx {
					firstIdx = idx
					first = v
				}
			}
			if endIdx > -1 {
				if endIdx >= lastIdx {
					lastIdx = endIdx
					last = v
				}
			}
		}
		final, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(final)
		sum = sum + final

	}
	log.Printf("Sum is: %d", sum)
}
