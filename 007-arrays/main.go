package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read from stdin
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	// Handle the first line of input
	_, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	// Handle input array and split strings into array
	arr := strings.Split(readLine(reader), " ")

	// Reverse the array then print in space separated format
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println(strings.Join(arr, " "))
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
