package main

import (
	"fmt"
	"strconv"
	"strings"

	utils "github.com/username/advent-of-code/2024/gocode/util"
)

const SpecialCharacters string = "?~!@#$%^&*+=-'/<>;: []{}"

type CurrentMatch struct {
	value []rune
}

func main() {
	rawInput := utils.ReadFileString("input")

	// input := removeChars(string(rawInput), getSpecialCharacterSet())

	validInstructions := findValidInstructions(string(rawInput))

	multiplySum := reduce(validInstructions, 0, func(acc int, args []int, index int) int {
		return acc + (args[0] * args[1])
	})

	fmt.Println("Multiplication sum: ", multiplySum)
}

func reduce[ElT any, AccT any](iterable []ElT, startWith AccT, operator func(accumulator AccT, element ElT, index int) AccT) AccT {
	var accumulator AccT = startWith
	for i, el := range iterable {
		accumulator = operator(accumulator, el, i)
	}
	return accumulator
}

func findValidInstructions(input string) [][]int {
	instructions := [][]int{}
	// s := "bmul(34,4#53)mul(2,34)4)str(sf,f)bmul(34,)453"

	candidates := strings.Split(filterDonts(input), "mul(")
	for _, candidate := range candidates {
		args1 := strings.Split(candidate, ")")[0]
		args := strings.Split(args1, ",")

		if len(args) != 2 {
			continue
		}

		firstArg, err := strconv.Atoi(args[0])
		if err != nil {
			continue
		}
		secondArg, err2 := strconv.Atoi(args[1])
		if err2 != nil {
			continue
		}

		if len(args[0]) > 3 || len(args[0]) < 1 || len(args[1]) > 3 || len(args[1]) < 1 {
			continue
		}
		instructions = append(instructions, []int{firstArg, secondArg})
	}

	return instructions
}

func removeChars(s string, removalSet map[rune]struct{}) string {
	var newStringArray []rune
	for _, r := range s {
		if _, exists := removalSet[r]; !exists {
			newStringArray = append(newStringArray, r)
		}
	}
	return string(newStringArray)
}

func getSpecialCharacterSet() map[rune]struct{} {
	set := make(map[rune]struct{})
	for _, r := range SpecialCharacters {
		set[r] = struct{}{}
	}
	return set
}

func filterDonts(input string) string {
	var cleanDos []string
	for _, do := range strings.Split(input, "do()") {
		cleanDo := strings.Split(do, "don't()")[0]
		cleanDos = append(cleanDos, cleanDo)
	}
	var cleanDosStringBuilder strings.Builder
	for _, cleanDo := range cleanDos {
		cleanDosStringBuilder.WriteString(cleanDo)
	}
	return cleanDosStringBuilder.String()
}
