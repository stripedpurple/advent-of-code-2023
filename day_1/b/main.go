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
	dict := map[string][]byte{
		"one":   []byte("1"),
		"two":   []byte("2"),
		"three": []byte("3"),
		"four":  []byte("4"),
		"five":  []byte("5"),
		"six":   []byte("6"),
		"seven": []byte("7"),
		"eight": []byte("8"),
		"nine":  []byte("9"),
	}

	inputFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var total int32 = 0
	for scanner.Scan() {
		re := regexp.MustCompile(`one|two|three|four|five|six|seven|eight|nine`)
		text := re.ReplaceAllFunc([]byte(scanner.Text()), func(s []byte) []byte {
			return dict[string(s)]
		})
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
