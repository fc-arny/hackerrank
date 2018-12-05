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
func miniMaxSum(arr []int) {
	sort.Ints(arr)

	var min, max int

	for i := 0; i < length; i++ {
		min += arr[i]
	}

	for i := len(arr) - 1; i >= len(arr)-length; i-- {
		max += arr[i]
	}

	fmt.Printf("%v %v\n", min, max)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int(arrItemTemp)
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
