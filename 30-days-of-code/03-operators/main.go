// https://www.hackerrank.com/challenges/30-operators/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	cost, err := strconv.ParseFloat(readLine(reader), 64)
	_tip, err := strconv.ParseInt(readLine(reader), 10, 64)
	_tax, err := strconv.ParseInt(readLine(reader), 10, 64)

	if err != nil {
		panic(err)
	}

	tip := float64(_tip) / 100.0
	tax := float64(_tax) / 100.0

	res := cost + cost*tip + cost*tax

	fmt.Println(math.Round(res))
}

func readLine(r *bufio.Reader) string {
	val, _, err := r.ReadLine()

	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(val), "\r\n")
}
