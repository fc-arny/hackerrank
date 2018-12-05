// https://www.hackerrank.com/challenges/plus-minus/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	readLine := func(r *bufio.Reader) string {
		val, _, err := r.ReadLine()

		if err == io.EOF {
			return ""
		}

		return strings.TrimRight(string(val), "\r\n")
	}

	reader := bufio.NewReader(os.Stdin)
	cnt, _ := strconv.ParseFloat(readLine(reader), 64)
	arrStr := readLine(reader)
	var res = make(map[string]float64)

	arr := strings.Split(arrStr, " ")

	for _, v := range arr {
		iVal, _ := strconv.ParseFloat(strings.TrimSpace(v), 64)

		switch {
		case iVal > 0:
			res["pos"]++
		case iVal < 0:
			res["neg"]++
		default:
			res["zero"]++
		}
	}

	pos := res["pos"] / cnt
	neg := res["neg"] / cnt
	zero := res["zero"] / cnt
	fmt.Printf("%.6f\n%.6f\n%.6f\n", pos, neg, zero)
}
