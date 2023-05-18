package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Open("./in.txt")
	defer f.Close()
	checkErr(err)

	scanner := bufio.NewScanner(f)

	var calories []int
	var currentCalories int
	var mostCalories int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Println(currentCalories)
			calories = append(calories, currentCalories)
			currentCalories = 0
			continue
		}
		caloriesAmt, err := strconv.Atoi(line)
		checkErr(err)
		currentCalories += caloriesAmt
	}

	sort.Ints(calories)
	for _, c := range calories[len(calories)-3:] {
		fmt.Println(c)
		mostCalories += c
	}

	fmt.Printf("Most Calories %d\n", mostCalories)
}
