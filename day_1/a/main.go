package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var total int32 = 0
	for scanner.Scan() {

		text := scanner.Text()
		left := 0
		right := len(text) - 1
		firstLetter := ""
		lastLetter := ""

		for {
			if left > right {
				break
			}

			if left == right {
				lastLetter = firstLetter
			}

			if firstLetter == "" {
				first := string(text[left])
				_, err := strconv.Atoi(first)

				if err != nil {
					left++
					continue
				}
				firstLetter = first
			}

			if lastLetter == "" {
				last := string(text[right])
				_, err2 := strconv.Atoi(last)

				if err2 != nil {
					right--
					continue
				}

				lastLetter = last
			}

			newVal, _ := strconv.Atoi(firstLetter + lastLetter)

			total += int32(newVal)
			break

		}
	}
	fmt.Println(total)
}
