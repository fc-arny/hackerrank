// https://www.hackerrank.com/challenges/30-data-types/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	// Declare second integer, double, and String variables.
	// Read and save an integer, double, and String to your variables.

	var ii int64
	// Print the sum of both integer variables on a new line.
	if scanner.Scan() {
		ii, _ = strconv.ParseInt(scanner.Text(), 10, 64)

	}

	var dd float64
	// Print the sum of the double variables on a new line.
	if scanner.Scan() {
		dd, _ = strconv.ParseFloat(scanner.Text(), 64)

	}

	var ss string
	// Concatenate and print the String variables on a new line
	// The 's' variable above should be printed first.
	if scanner.Scan() {
		ss = scanner.Text()
	}

	fmt.Printf("%d\n", uint64(ii)+i)
	fmt.Printf("%.1f\n", dd+d)
	fmt.Printf("%s%s\n", s, ss)

}
