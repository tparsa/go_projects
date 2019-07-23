package main

import (
	"fmt"
	"math"
)

func prime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}

func main() {

	var n int
	fmt.Scan(&n)

	for i := 2; i <= n; i++ {
		if prime(i) {
			fmt.Println(i)
		}
	}
}