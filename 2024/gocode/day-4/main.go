package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Match struct {
	point      Point
	directionX int
	directionY int
}

func main() {
	// rawInput := utils.ReadFileString("input.txt")

	// matrix := buildMatrix(rawInput)
	matrix := [][]string{
		{".", ".", ".", "X", "X", "M", "A", "S", "."},
		{".", "S", "A", "M", "X", "M", "S", ".", ".", "."},
		{".", ".", ".", "S", ".", ".", "A", ".", ".", "."},
		{".", ".", "A", ".", "A", ".", "M", "S", ".", "X"},
		{"X", "M", "A", "S", "A", "M", "X", ".", "M", "M"},
		{"X", ".", ".", ".", ".", ".", "X", "A", ".", "A"},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "S"},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "A"},
		{".", ".", "M", ".", "M", ".", "M", ".", "M", "M"},
		{".", "X", ".", "X", ".", "X", "M", "A", "S", "X"},
	}

	count := countByWord(matrix, "XMAS")
	fmt.Println("XMAS count: ", count)
}

func countByWord(matrix [][]string, word string) int {
	startingPoints := getPointsByChar(matrix, string(word[0]))

	// find count for each starting point
	count := 0
	matches := []Match{}
	for _, point := range startingPoints {
		matchesForWordAtPoint := getMatchesForWordAtPoint(matrix, "XMAS", point)
		matches = append(matches, matchesForWordAtPoint...)
		count += len(matchesForWordAtPoint)
	}

	pointMatchFrequencies := make(map[Match]int)
	for _, match := range matches {
		pointMatchFrequencies[match]++
	}
	fmt.Println(pointMatchFrequencies)
	return count
}

func getMatchesForWordAtPoint(matrix [][]string, target string, point Point) []Match {
	matches := []Match{}
	directions := [][]int{ // x, y
		{-1, -1}, // up left
		{0, -1},  // up
		{1, -1},  // up right
		{1, 0},   // right
		{1, 1},   // down right
		{0, 1},   // down
		{-1, 1},  // down left
		{-1, 0},  // left
	}

	for _, direction := range directions {
		if testInDirection(matrix, target, point, direction) {
			matches = append(matches, Match{
				point:      point,
				directionX: direction[0],
				directionY: direction[1],
			})
		}
	}
	return matches
}

func testInDirection(matrix [][]string, target string, point Point, direction []int) bool {
	testWord := string(target[0])
	for multiplier := 1; multiplier < len(target); multiplier++ {
		nextPoint := Point{
			x: point.x + (multiplier * direction[0]),
			y: point.y + (multiplier * direction[1]),
		}
		if inBounds(len(matrix), len(matrix[0]), nextPoint) {
			testWord += string(matrix[nextPoint.y][nextPoint.x])
		}
	}
	return testWord == target
}

func inBounds(maxY int, maxX int, nextPoint Point) bool {
	if nextPoint.x < 0 || nextPoint.x >= maxX {
		return false
	}
	if nextPoint.y < 0 || nextPoint.y >= maxY {
		return false
	}
	return true
}

func getPointsByChar(matrix [][]string, targetChar string) []Point {
	points := []Point{}
	for y, line := range matrix {
		for x, char := range line {
			if char == targetChar {
				points = append(points, Point{x: x, y: y})
			}
		}
	}
	return points
}

func buildMatrix(input string) [][]string {
	matrix := [][]string{}

	line := []string{}
	for _, char := range input {
		if char == '\n' {
			matrix = append(matrix, line)
			line = []string{}
			continue
		}
		line = append(line, string(char))
	}
	return matrix
}
