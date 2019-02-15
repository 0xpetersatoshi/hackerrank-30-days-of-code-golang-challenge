package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the factorial function below.
func factorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}

func main() {
	// Reading from stdin
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// Create output file to store result
	stdout, err := os.Create("result.txt")
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	// Parse input in and run factorial function
	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := factorial(n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
