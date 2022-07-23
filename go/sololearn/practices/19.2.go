/* WORD REPEATER
* You need to make a program which repeats a given word N number of times.
*
* The program needs to take a string and a number as input and output the string repeated the givern number of times on separate lines.
*
* The given code takes the inputs and calls the repeat() function with those arguments.
* You need to create the repeat() function.
*
* Sample Input:
* hello
* 4
*
* Sample Output:
* hello
* hello
* hello
* hello
*
* ! Use a loop in the function to repeat the output.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

func repeat(word string, reps int) {
	for rep := 0; rep < reps; rep++ {
		fmt.Println(word)
	}
}

func main() {
	var w string
	fmt.Scanln(&w)
	var x int
	fmt.Scanln(&x)

	repeat(w, x)
}
