package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const tolerance = 100

type Day struct {
	Date  string
	Total int
}

func main() {
	sex := flag.String("sex", "male", "male | female")
	age := flag.Int("age", 0, "age in years")
	height := flag.Float64("height", 0, "height in cm")
	weight := flag.Float64("weight", 0, "weight in kg")
	activity := flag.Float64("activity", 1.2, "activity factor (1.2–1.9)")
	filePath := flag.String("file", "calories.txt", "path to calories file")

	flag.Parse()

	if *age == 0 || *height == 0 || *weight == 0 {
		panic("age, height and weight are required")
	}

	limit := calculateTDEE(*sex, *age, *height, *weight, *activity)
	fmt.Printf("Норма: %.0f ккал\n\n", limit)

	days := parseFile(*filePath)
	printSummary(days, int(limit))
}

func calculateTDEE(sex string, age int, height, weight, activity float64) float64 {
	var bmr float64

	if sex == "female" {
		bmr = 10*weight + 6.25*height - 5*float64(age) - 161
	} else {
		bmr = 10*weight + 6.25*height - 5*float64(age) + 5
	}

	return bmr * activity
}

func parseFile(path string) []Day {
	file, err := os.Open(path)
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

		if isDate(line) {
			currentDate = line
			continue
		}

		if strings.HasPrefix(line, "==") && currentDate != "" {
			raw := strings.TrimSpace(strings.TrimPrefix(line, "=="))
			total, err := strconv.Atoi(raw)
			if err != nil {
				continue
			}

			days = append(days, Day{
				Date:  currentDate,
				Total: total,
			})
		}
	}

	return days
}

func printSummary(days []Day, limit int) {
	for _, day := range days {
		switch {
		case day.Total > limit+tolerance:
			fmt.Printf("\033[31m%s = %d 🔴 выше нормы\033[0m\n", day.Date, day.Total)
		case day.Total >= limit-tolerance:
			fmt.Printf("\033[33m%s = %d 🟡 около нормы\033[0m\n", day.Date, day.Total)
		default:
			fmt.Printf("\033[32m%s = %d 🟢 ниже нормы\033[0m\n", day.Date, day.Total)
		}
	}
}

func isDate(s string) bool {
	return len(s) == 5 && s[2] == '.'
}
