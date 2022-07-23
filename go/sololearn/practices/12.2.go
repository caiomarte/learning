/* FEET TO INCHES CONVERTER
* You need to create a program that converts feet to inches and outputs the resulting value.
* The feet value is providd as the input to the program.
* 1 foot = 12 inches.
*
* Sample Input:
* 5
*
* Sample Output:
* 60
*
* ! The input is an integer.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

func main() {
	var f int
	fmt.Scanln(&f)

	i := f * 12
	fmt.Println(i)
}
