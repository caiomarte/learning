/* GO!
* You are given code that includes calls to a function calles "talk()", which should output the text "GO".
* The function is called 3 times in main(), but the code is missing the function declaration.
*
* Add the missing code to declare the talk() function, which should output "GO" to the screen.
*
* ! Note, that the output should be in all caps.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

func talk() {
	fmt.Println("GO")
}

func main() {
	talk()
	talk()
	talk()
}
