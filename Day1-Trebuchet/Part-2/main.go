package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

var NUMBER_MAP = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func findFirstMatchingDigit(str string) (string, bool) {
	if len(str) > 0 && unicode.IsDigit(rune(str[0])) {
		return string(str[0]), true
	}

	length := len(str)
	if length >= 3 {
		// Using min function to avoid out-of-bound errors
		for i := 3; i <= min(5, length); i++ {
			key := str[:i]
			if value, exists := NUMBER_MAP[key]; exists {
				return value, true
			}
		}
	}

	return "", false
}

func trebuchetSolver() int {
	file, err := os.Open("./Day1-Trebuchet/Part-2/input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineLength := len(line)
		firstDigit := ""
		lastDigit := ""
		for i, j := 0, lineLength-1; i < lineLength && j >= 0; {
			if value, exists := findFirstMatchingDigit(line[i:]); exists && firstDigit == "" {
				firstDigit = value
			}

			if value, exists := findFirstMatchingDigit(line[j:]); exists && lastDigit == "" {
				lastDigit = value
			}

			if firstDigit != "" && lastDigit != "" {
				break
			}

			i++
			j--
		}

		result, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatalf("error converting to integer: %s", err)
		}

		sum += result
	}

	return sum
}

func main() {
	sum := trebuchetSolver()
	fmt.Println(sum)
}
