package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	numRed   = 12
	numBlue  = 14
	numGreen = 13
	input    = "input.txt"
)

func main() {
	file, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		split := strings.Split(text, ":")
		game := split[0]
		selections := split[1]
		gameNum, err := strconv.Atoi(strings.Split(game, " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(selections, ";")
		valid := true
		for _, round := range rounds {
			colours := strings.Split(round, ",")
			for _, colour := range colours {
				countColour := strings.Split(strings.TrimSpace(colour), " ")
				count, err := strconv.Atoi(countColour[0])
				if err != nil {
					log.Fatal(err)
				}
				switch countColour[1] {
				case "red":
					if count > numRed {
						valid = false
					}
				case "blue":
					if count > numBlue {
						valid = false
					}
				case "green":
					if count > numGreen {
						valid = false
					}
				}
			}
		}
		if valid {
			sum = sum + gameNum
			log.Println(text)
			log.Println("----")

		}
	}
	log.Println(sum)
}
