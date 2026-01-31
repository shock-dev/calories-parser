package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const kcalPerKg = 7700.0

type Day struct {
	Date  string
	Total int
}

func main() {
	n := flag.Int("n", 0, "daily norm calories")
	flag.Parse()

	file, err := os.Open("calories.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var days []Day
	var currentDate string
	sumActual := 0

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
			days = append(days, Day{currentDate, dayTotal})
			sumActual += dayTotal
		}
	}

	for _, day := range days {
		fmt.Printf("%s = %d\n", day.Date, day.Total)
	}

	fmt.Println("----------------")

	daysCount := len(days)
	avg := float64(sumActual) / float64(daysCount)

	if *n == 0 {
		fmt.Printf("Итого: %d ккал\n", sumActual)
		fmt.Printf("Среднее: %.0f ккал/день\n", avg)
		return
	}

	sumNorm := (*n) * daysCount
	burned := sumNorm - sumActual

	fmt.Printf("n: %d ккал\n\n", *n)
	fmt.Printf("Норма за период: %d ккал\n", sumNorm)
	fmt.Printf("У меня:          %d ккал\n\n", sumActual)
	fmt.Printf("Сожжено: %d ккал\n", burned)
	fmt.Printf("≈ %.2f кг\n", float64(burned)/kcalPerKg)
	fmt.Printf("Среднее потребление: %.0f ккал/день\n", avg)
}

func isDate(s string) bool {
	return len(s) == 5 && s[2] == '.'
}
