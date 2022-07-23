/* ROUTE MAKER
* You are making a program for a navigation system.
* Your program needs to take a number N as input, followed by N strings, which represent city names.
*
* Store the city names in a slice.
* Declare a variadic function route(), which takes its arguments from the slice.
* The route() function should output the route, which combines the city names using an arrow.
* Check the sample input/output for reference.
*
* Sample Input:
* 3
* Boston
* Chicago
* Washington
*
* Sample Output:
* Boston->Chicago->Washington->
*
* ! Create the output by iterating over the aruments and concatenating an arrow after each value.
*
* Source: https://www.sololearn.com/learning/1164 */

package main

import "fmt"

//create the route() function
func route(cities ...string) {
	for _, c := range cities {
		fmt.Printf("%s->", c)
	}
}

func main() {
    var n int
    fmt.Scanln(&n)

    var cities []string
    //take n strings as input and append them to the slice
	cities = make([]string, 0)
	var input string
	for i := 0; i < n; i++ {
		fmt.Scanln(&input)
		cities = append(cities, input)
	}

    //call the route() function
    route(cities...)
}