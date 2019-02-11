package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the solve function below.
func solve(mealCost float64, tipPercent int32, taxPercent int32) {
	tip := mealCost * (float64(tipPercent) / 100)
	tax := mealCost * (float64(taxPercent) / 100)
	totalCost := math.Round(mealCost + tip + tax)

	fmt.Println(totalCost)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	mealCost, err := strconv.ParseFloat(readLine(reader), 64)
	checkError(err)

	tipPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	tipPercent := int32(tipPercentTemp)

	taxPercentTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	taxPercent := int32(taxPercentTemp)

	solve(mealCost, tipPercent, taxPercent)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
