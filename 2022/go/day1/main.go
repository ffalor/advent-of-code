package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ffalor/advent-of-code/util"
)

//go:embed in.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")

	if len(input) == 0 {
		panic("Empty in.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "Problem part to run")
	flag.Parse()

	fmt.Printf("Running part: %d", part)

	if part == 1 {
		part1()
	} else {
		part2()
	}
}

func part1() {
	mostCalories := 0
	currentCalories := 0

	f, err := os.Open("./in.txt")
	defer f.Close()
	util.ErrPanic(err)

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
		util.ErrPanic(err)

		currentCalories += calories
	}

	fmt.Printf("Most Calories %d\n", mostCalories)
}

func part2() {
	f, err := os.Open("./in.txt")
	defer f.Close()
	util.ErrPanic(err)

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
		util.ErrPanic(err)
		currentCalories += caloriesAmt
	}

	sort.Ints(calories)
	for _, c := range calories[len(calories)-3:] {
		fmt.Println(c)
		mostCalories += c
	}

	fmt.Printf("Most Calories %d\n", mostCalories)
}
