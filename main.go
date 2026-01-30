package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Day struct {
	Date  string
	Total int
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
	totalAllDays := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if isDate(line) {
			currentDate = line
			continue
		}

		if strings.HasPrefix(line, "==") && currentDate != "" {
			raw := strings.TrimSpace(strings.TrimPrefix(line, "=="))
			dayTotal, err := strconv.Atoi(raw)
			if err != nil {
				continue
			}

			days = append(days, Day{
				Date:  currentDate,
				Total: dayTotal,
			})

			totalAllDays += dayTotal
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for _, day := range days {
		fmt.Printf("%s = %d\n", day.Date, day.Total)
	}

	fmt.Println("----------------")
	fmt.Printf("Итого: %d ккал\n", totalAllDays)
}

func isDate(s string) bool {
	return len(s) == 5 && s[2] == '.'
}
