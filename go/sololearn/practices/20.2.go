/* SQUARE
* You need to create a function which takes a number as its argument and retunrs its square.
*
* Your program needs to take the number from input and pass it to the function, which should return the result.
*
* ! The main function already includes the code to take a number as input and call the square() function.
*   Complete the program by creating the function.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

func square(number int) int {
	return number * number
}

func main() {
	var x int
	fmt.Scanln(&x)

	res := square(x)
	fmt.Println(res)
}
