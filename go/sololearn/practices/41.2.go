/* CONCURRENT COUNTER
* You are given an aarray of numbers.
* You need to take a number as input and output the number of times that number occurs in the array.
* To make your program run faster, you decide to split the array into 2 halves and run 2 Goroutines for each half.
* After each Goroutine is done, add together the results for each half and output the final result.
*
* The given program takes a number as input and calls the count() function as two Goroutines for each half of the data array.
*
* Task:
* 1. Define the count() function, which should take a number, an array, and a channel as arguments.
*    The function should then calculate the number of times the given number occurs in the array and send the result to the channel.
*
* 2. Collect the results of the 2 Goroutines and output their sum.
*
* ! Take a look ad the Goroutines in the code to understand how the count() function should work.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

//define the count() function
func count(number int, array []int, channel chan int) {
    ocurrency := 0
    for _, i := range array {
        if i == number {
            ocurrency++
        }
    }
    channel <- ocurrency
}

func main() {
    data := []int{12, 45, 88, 42, 0, 98, 102, 42, 77, 42, 1, 8, 7, 55, 4, 12, 87, 90, 42, 42, 11, 2, 6, 53, 90, 100, 4, 32, 8}
    
    var num int
    fmt.Scanln(&num)

    ch1 := make(chan int)
    ch2 := make(chan int)

    go count(num, data[:len(data)/2], ch1)
    go count(num, data[len(data)/2:], ch2)

    //output the final result here
    fmt.Println(<-ch1 + <-ch2)
    
}