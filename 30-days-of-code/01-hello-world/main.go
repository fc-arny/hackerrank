// https://www.hackerrank.com/challenges/30-hello-world/problem
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT

	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Hello, World")
	for s.Scan() {
		fmt.Printf(s.Text())
	}

	

}

// func readline() string {
// 	reader := bufio.NewReader(os.Stdin)
// 	str, _, err := reader.ReadLine()

// 	if err == io.EOF {
// 		return ""
// 	}

// 	return string(str)
// }
