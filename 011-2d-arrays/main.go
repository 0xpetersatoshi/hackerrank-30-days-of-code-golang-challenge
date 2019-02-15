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

	// Create a 6 x 6 array
	n := 6
	arr := make([][]float64, n)
	for inner := range arr {
		arr[inner] = make([]float64, n)
	}

	// Iterate through input data and insert parsed ints into 2d array
	for i := 0; i < n; i++ {
		// Read line and split into array of elements
		strArr := strings.Split(readLine(reader), " ")
		for j := 0; j < n; j++ {
			number, _ := strconv.ParseFloat(strArr[j], 64)
			arr[i][j] = number
		}
	}

	// Calculate hourglass sum and max value
	// Set max to lowest possible hourglass sum
	var max float64
	max = (-9) * 7

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr)-2; j++ {
			top := arr[i][j] + arr[i][j+1] + arr[i][j+2]
			middle := arr[i+1][j+1]
			bottom := arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			hourglassSum := top + middle + bottom

			max = math.Max(max, hourglassSum)
		}
	}

	fmt.Println(max)
}

func readLine(r *bufio.Reader) string {
	str, _, err := r.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.Trim(string(str), "\n")
}
