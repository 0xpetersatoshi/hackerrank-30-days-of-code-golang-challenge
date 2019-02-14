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

func solve(mealCost float64, tipPercent int32, taxPercent int32) {
	// Calculate tip, tax, and total cost
	tip := mealCost * (float64(tipPercent) / 100)
	tax := mealCost * (float64(taxPercent) / 100)
	totalCost := math.Round(mealCost + tip + tax)

	fmt.Println(totalCost)
}

func main() {
	// Reading from stdin
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// Parsing incoming strings to numbers and
	// calculating total meal cost
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

// Helper function to read incoming data
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

// Helper function to check for errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
