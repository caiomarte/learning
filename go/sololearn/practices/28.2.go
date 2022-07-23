/* RESET THE TIMER
* The given code declares a Timer struct, which is initialized in main().
* You need to add a reset functionality to the program.
*
* The reset() function needs to take a pointer to the struct as its argument and reset the seconds value to 0.
*
* ! Do not change the code in main(). Add the reset() function, so that the code you are given work as expected.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

type Timer struct {
	id      string
	seconds int
}

func reset(timer *Timer) {
	timer.seconds = 0
}

func main() {
	var x int
	fmt.Scanln(&x)
	t := Timer{"ABC", x}

	reset(&t)
	fmt.Println(t)
}
