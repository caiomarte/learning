/* SAY THE NUMBERS
* You are making a robot that can speak numbers.
* Your robot should take 3 numbers in the range of 0-10 as input and output the corresponding texts in English.
*
* Sample Input:
* 8
* 0
* 5
*
* Sample Output:
* Eight
* Zero
* Five
*
* Source: https://www.sololearn.com/learning/eom-project/1164/1188 */

package main

import "fmt"

func main() {
	var numbers [3]int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&numbers[i])

		if numbers[i] >= 0 && numbers[i] <= 10 {
			fmt.Println(numberToText(numbers[i]))
		} else {
			fmt.Println("Value not allowed! Please enter a whole number from 0 to 10.")
		}
	}
}

func numberToText(number int) string {
	var text string
	switch number {
	case 0:
		text = "Zero"
	case 1:
		text = "One"
	case 2:
		text = "Two"
	case 3:
		text = "Three"
	case 4:
		text = "Four"
	case 5:
		text = "Five"
	case 6:
		text = "Six"
	case 7:
		text = "Seven"
	case 8:
		text = "Eight"
	case 9:
		text = "Nine"
	case 10:
		text = "Ten"
	}
	return text
}
