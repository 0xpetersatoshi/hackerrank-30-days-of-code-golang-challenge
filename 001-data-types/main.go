package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	// Reading input from stdin
	scanner := bufio.NewScanner(os.Stdin)

	var incomingInt int64
	var incomingFloat float64
	var incomingString string
	inputVals := []string{}

	// Append incoming values to inputVals until
	// there is no more input to read
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}
		inputVals = append(inputVals, input)
	}

	// Convert inputs to appropriate types, add/concatenate
	// with starter variables, and print results
	incomingInt, _ = strconv.ParseInt(inputVals[0], 10, 64)
	incomingFloat, _ = strconv.ParseFloat(inputVals[1], 64)
	incomingString = inputVals[2]

	fmt.Println(i + uint64(incomingInt))
	fmt.Printf("%.1f\n", d+incomingFloat)
	fmt.Println(s + incomingString)
}
