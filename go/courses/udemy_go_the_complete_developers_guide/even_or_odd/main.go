package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := []int{0,1,2,3,4,5,6,7,8,9,10}

	for _, integer := range n {
		if integer % 2 == 0 {
			fmt.Println(strconv.Itoa(integer) + " is even")
		} else {
			fmt.Println(strconv.Itoa(integer) + " is odd")
		}
	}
}

// 1. Briefly explain your approach and possible challenges you overcame.
// I used a for loop and the modulo operator to iterate over the slice of integers and print them to the command line.
// Two challenges I overcame were finding the package to turn an integer into a string and initializing a slice with values