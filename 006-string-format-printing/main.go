package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}

	s = strings.Trim(s, "\n")

	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	evens := []string{}
	odds := []string{}
	results := []string{}
	var spaced string
	var e string
	var o string

	for i := 0; i < int(n); i++ {
		s, _ := reader.ReadString('\n')
		s = strings.Trim(s, "\n")

		for idx, val := range s {
			if idx%2 == 0 {
				evens = append(evens, string(val))
			} else {
				odds = append(odds, string(val))
			}
		}
		e = strings.Join(evens, "")
		o = strings.Join(odds, "")
		tmp := []string{e, o}
		spaced = strings.Join(tmp, " ")
		results = append(results, spaced)
		evens = []string{}
		odds = []string{}
	}

	result := strings.Join(results, "\n")
	fmt.Println(result)
}
