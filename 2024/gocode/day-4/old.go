package main

// import (
// 	"fmt"

// 	utils "github.com/username/advent-of-code/2024/gocode/utils"
// )

// type Point struct {
// 	x int
// 	y int
// }

// type Match struct {
// 	wordTail     string
// 	currentPoint Point
// 	seen         map[Point]struct{}
// }

// func main() {
// 	// rawInput := utils.ReadFileString("input.txt")

// 	// matrix := buildMatrix(rawInput)
// 	matrix := [][]string{
// 		{"S", "A", "M", "X"},
// 		{"M", "M", "M", "S"},
// 		{"A", "A", "A", "S"},
// 		{"S", "M", "A", "S"},
// 	}

// 	count := countByWord(matrix, "XMAS")

// 	fmt.Println(matrix)
// 	fmt.Println("XMAS count: ", count)
// }

// func getPointsByChar(matrix [][]string, targetChar string) []Point {
// 	points := []Point{}
// 	for y, line := range matrix {
// 		for x, char := range line {
// 			if char == targetChar {
// 				points = append(points, Point{x: x, y: y})
// 			}
// 		}
// 	}
// 	return points
// }

// func countByWord(matrix [][]string, word string) int {
// 	startingPoints := getPointsByChar(matrix, "X")

// 	// Add all starting points to queue as new empty matches
// 	queue := utils.NewQueue()
// 	for _, point := range startingPoints {
// 		queue.Enqueue(Match{
// 			// Start with the first letter at point
// 			wordTail:     matrix[point.y][point.x],
// 			currentPoint: point,
// 			seen:         map[Point]struct{}{},
// 		})
// 	}

// 	count := 0
// 	for !queue.IsEmpty() {
// 		currentMatch := queue.Dequeue().(Match)

// 		if currentMatch.wordTail == word {
// 			count++
// 			continue
// 		}

// 		// Add all eligible neighbor points to queue
// 		neighborPoints := getNeighborPoints(matrix, currentMatch.currentPoint)
// 		for _, next := range neighborPoints {
// 			// Do not backtrack to a point already explored
// 			if _, exists := currentMatch.seen[next]; exists {
// 				continue
// 			}
// 			newMatch := Match{
// 				wordTail:     currentMatch.wordTail + matrix[next.y][next.x],
// 				currentPoint: next,
// 				seen:         currentMatch.seen,
// 			}
// 			newMatch.seen[currentMatch.currentPoint] = struct{}{}
// 			queue.Enqueue(newMatch)
// 		}
// 	}
// 	return count
// }

// func getNeighborPoints(matrix [][]string, currentPoint Point) []Point {
// 	pointsToExplore := make([]Point, 0, 10)
// 	maxX := len(matrix[0]) - 1
// 	maxY := len(matrix) - 1
// 	// Add in all directions, if within bounds
// 	// up and left
// 	if currentPoint.x-1 >= 0 && currentPoint.y-1 >= 0 {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x - 1,
// 			y: currentPoint.y - 1,
// 		})
// 	}
// 	// up
// 	if currentPoint.y-1 >= 0 {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x,
// 			y: currentPoint.y - 1,
// 		})
// 	}
// 	// up and right
// 	if currentPoint.x+1 <= maxX && currentPoint.y-1 >= 0 {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x + 1,
// 			y: currentPoint.y - 1,
// 		})
// 	}
// 	// right
// 	if currentPoint.x+1 <= maxX {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x + 1,
// 			y: currentPoint.y,
// 		})
// 	}
// 	// down and right
// 	if currentPoint.x+1 <= maxX && currentPoint.y+1 <= maxY {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x + 1,
// 			y: currentPoint.y + 1,
// 		})
// 	}
// 	// down
// 	if currentPoint.y+1 <= maxY {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x,
// 			y: currentPoint.y + 1,
// 		})
// 	}
// 	// down and left
// 	if currentPoint.x-1 >= 0 && currentPoint.y+1 <= maxY {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x - 1,
// 			y: currentPoint.y + 1,
// 		})
// 	}
// 	// left
// 	if currentPoint.x-1 >= 0 {
// 		pointsToExplore = append(pointsToExplore, Point{
// 			x: currentPoint.x - 1,
// 			y: currentPoint.y,
// 		})
// 	}

// 	return pointsToExplore
// }

// func buildMatrix(input []string) [][]string {
// 	matrix := [][]string{}

// 	line := []string{}
// 	for _, char := range input {
// 		if char == "\n" {
// 			matrix = append(matrix, line)
// 			line = []string{}
// 			continue
// 		}
// 		line = append(line, char)
// 	}
// 	return matrix
// }
