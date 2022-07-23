/* TEMPERATURE CHECKER
* You are writing a program for a temperature checking system at an airport.
* The system measures the body temperature of a person and needs to output "Allowed" if it is in the normal range, or "Fever" if it is higher than normal.
* Normal: up to 99.5 °F
* Fever: > 99.5 °F
*
* The program should take the temperature as a float from input and output the corresponding message.
* Sample Input:
* 101.3
*
* Sample Output:
* Fever
*
* ! Use an if/else statement to make the decision.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

func main() {
	var temperature float32
	fmt.Scanln(&temperature)

	if temperature > 99.5 {
		fmt.Println("Fever")
	} else {
		fmt.Println("Allowed")
	}
}
