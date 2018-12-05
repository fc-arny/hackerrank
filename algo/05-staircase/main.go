// https://www.hackerrank.com/challenges/staircase/problem

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the staircase function below.
func staircase(n int32) {
	j := n

	for i := 1; i <= int(n); i++ {
		j--
		var sc []string
		for k := 0; k < int(j); k++ {
			sc = append(sc, " ")
		}

		for k := 0; k < int(i); k++ {
			sc = append(sc, "#")
		}

		fmt.Println(strings.Join(sc, ""))
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	staircase(n)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
