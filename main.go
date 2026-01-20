package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Day struct {
	Date  string
	Total string
}

func main() {
	file, err := os.Open("calories.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var days []Day
	var currentDate string

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		// дата вида 08.01
		if isDate(line) {
			currentDate = line
			continue
		}

		// итог вида "== 1940"
		if strings.HasPrefix(line, "==") && currentDate != "" {
			total := strings.TrimSpace(strings.TrimPrefix(line, "=="))
			days = append(days, Day{
				Date:  currentDate,
				Total: total,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// вывод
	for _, day := range days {
		fmt.Printf("%s = %s\n", day.Date, day.Total)
	}
}

func isDate(s string) bool {
	// формат XX.XX
	if len(s) != 5 {
		return false
	}
	return s[2] == '.'
}
