package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const (
	input = "input.txt"
)

func main() {
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
		last := ""
		log.Println(text)
		for _, c := range text {
			// c will be the unicode
			if unicode.IsNumber(c) {
				num := fmt.Sprintf("%c", c)
				if first == "" {
					first = num
				}
				last = num
			}
		}
		final, err := strconv.Atoi(first + last)
		if err != nil {
			log.Fatal(err)
		}
		sum = sum + final

	}
	log.Println(sum)
}
