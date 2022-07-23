/* AGE ON MARS
* How old would you be on Mars?
* A year on Earth has 365 days, while a year on Mars has 687 days.
* Create a program that takes your age in Earth years as input, and outputs the corresponding age on Mars.
* The given program takes an integer as input and passes it to the mars_age() function as argument.
* Complete the function to calculate and return the corresponding Mars age.
*
* Sample Input:
* 42
*
* Sample Output:
* 22
*
* Source: https://www.sololearn.com/learning/eom-project/1164/1189 */

package main

import "fmt"

func main() {
	var age int
	fmt.Scanln(&age)

	mars := mars_age(age)
	fmt.Println(mars)
}

func mars_age(ageOnEarth int) int {
	return ageOnEarth * 365 / 687
}
