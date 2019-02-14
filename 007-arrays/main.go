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
	reader := bufio.NewReader(os.Stdin)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	n := int(nTemp)

	strArr := strings.Split(readLine(reader), " ")

	reversedArr := []string{}

	for i := n - 1; i >= 0; i-- {
		reversedArr = append(reversedArr, strArr[i])
	}

	output := strings.Join(reversedArr, " ")
	fmt.Println(output)

}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
