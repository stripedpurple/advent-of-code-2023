package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	dict := map[string]string{
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

	inputFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var total int32 = 0
	for scanner.Scan() {
		re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine|[1-9]`)
		firstLetter := ""
		lastLetter := ""
		text := scanner.Text()

		for i := 0; i < len(text); i++ {
			match := re.FindString(text[i:])

			if match == "" {
				continue
			}

			if firstLetter == "" {
				if len(match) > 0 {
					firstLetter = dict[match]
					lastLetter = dict[match]
					continue
				}
			}

			if len(match) > 0 {
				lastLetter = dict[match]
			}
		}

		newVal, _ := strconv.Atoi(firstLetter + lastLetter)
		total += int32(newVal)
	}
	fmt.Println(total)
}
