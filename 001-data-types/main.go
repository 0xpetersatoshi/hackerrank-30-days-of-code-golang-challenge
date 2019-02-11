package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	// var i uint64 = 4
	// var d float64 = 4.0
	// var s string = "HackerRank "

	logstring :=
		`12
4.0
is the best place to learn and practice coding!`

	scanner := bufio.NewScanner(strings.NewReader(logstring))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		input := scanner.Text()

		fmt.Println(input)
	}
}
