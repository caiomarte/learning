/* MATCH RESULTS
* You are making a program to analyze sport match results and calculate the points of the given team.
* The match results are stored in an array called results.
* Each match has one of the following results:
* - "w" - won
* - "l" - lost
* - "d" - draw
*
* A win adds three points, a draw adds one point, and a lost match does not add any points.
*
* Your program needs to take the last match result as input and append it to the results array.
* Afther that, calculate and output the total points the team gained from the results.
*
* Source: https://www.sololearn.com/learning/eom-project/1164/1191 */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}

	// Read last match result
	var lastMatchResult string
	fmt.Scanln(&lastMatchResult)

	// Append last match result to results array
	results = append(results, lastMatchResult)

	points, totalPoints := 0, 0
	for i, r := range results {
		// Replacing match result with corresponding points
		switch r {
		case "w":
			results[i] = "3"
		case "d":
			results[i] = "1"
		case "l":
			results[i] = "0"
		}

		// Converting results string value into integer
		points, _ = strconv.Atoi(results[i])
		// Summing points along iteration
		totalPoints += points
	}

	// Outputting total points
	fmt.Println(totalPoints)

}
