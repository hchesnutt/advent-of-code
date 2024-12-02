package main

import (
	"fmt"
	"math"
	"os"
	"strings"

	utils "github.com/username/advent-of-code/2024/gocode/util"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		fmt.Println("error reading file: ", err)
		return
	}

	reports := parseReports(string(input))

	var safeReportCount = calculateSafeReportCount(reports)
	fmt.Println("Safe report count: ", safeReportCount)
}

func parseReports(input string) [][]int {
	lines := strings.Split(input, "\n")

	var parsedLines [][]int
	for _, line := range lines {
		var parsedLine []int
		for _, char := range strings.Split(line, " ") {
			parsedLine = append(parsedLine, utils.ParseInt(char))
		}
		parsedLines = append(parsedLines, parsedLine)
	}

	return parsedLines
}

func calculateSafeReportCount(reports [][]int) int {
	var safeReportCount int
	for _, report := range reports {
		if isMonotonicWithinStepThreshholds(report, 1, 3) {
			safeReportCount++
		} else {
			for i := 0; i < len(report); i++ {
				reportWithRemoval := append([]int{}, report[0:i]...)
				reportWithRemoval = append(reportWithRemoval, report[i+1:]...)
				if isMonotonicWithinStepThreshholds(reportWithRemoval, 1, 3) {
					safeReportCount++
					break
				}
			}
		}

	}

	return safeReportCount
}

func isMonotonicWithinStepThreshholds(report []int, minStep int, maxStep int) bool {
	// Not monotonically increasing or decreasing
	if report[0] == report[1] {
		return false
	}
	monotonicallyIncreasing := report[0] < report[1]

	i := 0
	for i < len(report)-1 {
		curr := report[i]
		next := report[i+1]
		difference := next - curr
		distance := int(math.Abs(float64(difference)))

		// Check step constraint
		if distance < minStep || distance > maxStep {
			return false
		}

		// Check Monotonicity constraint
		if (monotonicallyIncreasing && difference < 0) || (!monotonicallyIncreasing && difference > 0) {
			return false
		}

		i++
	}

	return true
}
