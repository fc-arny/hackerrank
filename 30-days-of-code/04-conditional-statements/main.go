// https://www.hackerrank.com/challenges/30-conditional-statements/problem
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	n := readLine(r)
	fmt.Println(isWeird(n))
}

func isWeird(n int) string {
	var res string

	if n%2 != 0 {
		res = "Weird"
	} else {
		switch {
		case n >= 2 && n <= 5:
			res = "Not Weird"
		case n >= 6 && n <= 20:
			res = "Weird"
		default:
			res = "Not Weird"
		}
	}
	return res
}

func readLine(r *bufio.Reader) int {
	str, _, _ := r.ReadLine()
	s := string(str)

	v, _ := strconv.Atoi(s)
	return v
}
