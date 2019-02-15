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

func main() {
	// Read from stdin
	reader := bufio.NewReader(os.Stdin)
	input := readLine(reader)

	// Parse input into int
	n, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	// Convert/format int as string representation of binary number
	binary := strconv.FormatInt(n, 2)

	// Split binary string into array of 1s to determine number of consecutive 1s
	binaryArr := strings.Split(binary, "0")
	var max float64

	for _, val := range binaryArr {
		num, _ := strconv.ParseFloat(val, 64)
		max = math.Max(num, max)
	}

	// Output number of consecutive 1s
	consecutiveOnes := strconv.Itoa(int(max))
	fmt.Println(len(consecutiveOnes))

}

// Helper func to parse incoming input
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
