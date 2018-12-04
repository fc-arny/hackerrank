// https://www.hackerrank.com/challenges/mini-max-sum/problem
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const length = 4

// Complete the miniMaxSum function below.
func miniMaxSum(arr []int32) {
	var ints []int

	for _, v := range arr {
		ints = append(ints, int(v))
	}
	sort.Ints(ints)

	var min, max int

	for i := 0; i < length; i++ {
		min += ints[i]
	}

	for i := len(arr) - 1; i >= len(arr)-length; i-- {
		max += ints[i]
	}

	fmt.Printf("%v %v\n", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	miniMaxSum(arr)
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
