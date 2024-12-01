package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("../input-1")
	if err != nil {
		fmt.Println("error reading file")
		return
	}

	lefts, rights := parseToLeftsAndRights(string(input))

	if len(lefts) != len(rights) {
		panic("Oh no! Lefts and rights had different lengths")
	}

	slices.Sort(lefts)
	slices.Sort(rights)

	sum := sumDifferences(lefts, rights)
	fmt.Println("Sum of differences: ", sum)

	similarityScore := calculateSimilarityScore(lefts, rights)
	fmt.Println("Similarity score: ", similarityScore)
}

func parseInt(s string) int {
	var integer, err = strconv.Atoi(s)
	if err != nil {
		panic(s)
	}

	return integer
}

func parseToLeftsAndRights(input string) ([]int, []int) {
	var lines = strings.Split(input, "\n")

	var lefts []int
	var rights []int
	for _, line := range lines {
		tuple := strings.Split(line, "   ")

		lefts = append(lefts, parseInt(tuple[0]))
		rights = append(rights, parseInt(tuple[1]))
	}

	return lefts, rights
}

func sumDifferences(lefts []int, rights []int) int {
	var sum = 0
	for i := 0; i < len(lefts); i++ {
		sum += int(math.Abs(float64(lefts[i] - rights[i])))
	}
	return sum
}

func calculateSimilarityScore(lefts []int, rights []int) int {
	rightFrequencies := calcFrequencies(rights)

	var score = 0
	for _, left := range lefts {
		score += left * rightFrequencies[left]
	}
	return score
}

func calcFrequencies(arr []int) map[int]int {
	frequencies := make(map[int]int)

	for _, el := range arr {
		frequencies[el]++
	}

	return frequencies
}
