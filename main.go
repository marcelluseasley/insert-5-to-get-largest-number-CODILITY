package main

import (
	"fmt"
	"math"
	"strconv"
)


func main() {
	fmt.Println(Solution(30))

}


func Solution(N int) int {
	// write your code in Go 1.4
	digit := 5
	maximumNumber := -8000
	
	if N == 0 {
		return digit * 10
	}

	sign := float64(N) / math.Abs(float64(N))

	N = int(math.Abs(float64(N)))

	// number of digits in number
	sN := strconv.FormatInt(int64(N),10)
	digitPos := 1
	for i := 0; i <= len(sN); i++ {
		number := ((N / digitPos) * (digitPos * 10)) + (digit * digitPos) + (N % digitPos)
		fmt.Println(number)
		if number * int(sign) > maximumNumber {
			maximumNumber = int(sign) * number
		}
		digitPos *= 10
	}


	
	return maximumNumber
}
