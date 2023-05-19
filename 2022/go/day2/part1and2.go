package main

import (
	"bufio"
	"fmt"
	"os"
)

var partOneTable = map[string]int{
	"A X": 4, "A Y": 8, "A Z": 3,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 7, "C Y": 2, "C Z": 6,
}

var partTwoTable = map[string]int{
	"A X": 3, "A Y": 4, "A Z": 8,
	"B X": 1, "B Y": 5, "B Z": 9,
	"C X": 2, "C Y": 6, "C Z": 7,
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open("./in.txt")
	checkErr(err)

	var result int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		result += partTwoTable[scanner.Text()]
	}
	fmt.Println(result)
}
