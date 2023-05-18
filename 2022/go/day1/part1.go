package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	mostCalories := 0
	currentCalories := 0

	f, err := os.Open("./in.txt")
	defer f.Close()
	checkErr(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		calories, err := strconv.Atoi(line)
		checkErr(err)

		currentCalories += calories
	}

	fmt.Printf("Most Calories %d\n", mostCalories)
}
