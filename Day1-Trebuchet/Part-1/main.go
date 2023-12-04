package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func trebuchetSolver() int {
	file, err := os.Open("./Day1-Trebuchet/Part-1/input.txt")
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
			if unicode.IsDigit(rune(line[i])) && firstDigit == "" {
				firstDigit = string(line[i])
			}

			if unicode.IsDigit(rune(line[j])) && lastDigit == "" {
				lastDigit = string(line[j])
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
