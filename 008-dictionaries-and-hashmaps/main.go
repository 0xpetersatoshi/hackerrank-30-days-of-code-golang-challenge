package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	inputs := []string{}

	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			break
		}

		inputs = append(inputs, input)
	}

	nString := inputs[0]
	n, err := strconv.ParseInt(nString, 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	nameAndNums := inputs[1 : n+1]
	namesToCheck := inputs[n+1:]
	m := make(map[string]int64, n)

	for _, v := range nameAndNums {
		strArr := strings.Split(v, " ")
		name := strArr[0]
		nums := strArr[1]
		number, err := strconv.ParseInt(nums, 10, 64)
		if err != nil {
			fmt.Println(err)
		}

		m[name] = number
	}

	for _, name := range namesToCheck {
		num, ok := m[name]
		if !ok {
			fmt.Println("Not found")
		} else {
			fmt.Printf("%v=%v\n", name, num)
		}
	}
}
