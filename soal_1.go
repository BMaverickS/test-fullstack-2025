package main

import (
	"fmt"
	"math"
)

// soal nomor 1
func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res = res * i
	}
	res = int(math.Ceil(float64(res) / math.Pow(2.0, float64(n))))
	return res
}

func main() {
  // ubah angka sesuai kebutuhan
	fmt.Println(factorial(5))
}
