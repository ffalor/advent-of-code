package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func charToInt(c rune) int {
	if int(c) > 90 {
		return int(c) - 96
	} else {
		return int(c) - 38
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("./in.txt")
	checkErr(err)

	scanner := bufio.NewScanner(f)
	prioritySum := 0

	for scanner.Scan() {
		bp := scanner.Text()

		compartmentOne := bp[len(bp)/2:]
		compartmentTwo := bp[:len(bp)/2]

		for _, s := range compartmentOne {
			if strings.Contains(compartmentTwo, string(s)) {
				prioritySum += charToInt(s)
				break
			}
		}
	}

	fmt.Println(prioritySum)
}
