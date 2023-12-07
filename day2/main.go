package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
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

		split := strings.Split(text, ":")

		selections := split[1]
		if err != nil {
			log.Fatal(err)
		}

		rounds := strings.Split(selections, ";")
		maxRed := 0
		maxBlue := 0
		maxGreen := 0
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
					if count > maxRed {
						maxRed = count
					}
				case "blue":
					if count > maxBlue {
						maxBlue = count
					}
				case "green":
					if count > maxGreen {
						maxGreen = count
					}
				}
			}
		}
		sum = sum + (maxBlue * maxRed * maxGreen)
	}
	log.Println(sum)
}
