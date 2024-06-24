package main

import (
	"fmt"
	"math"
)

const s string = "CONSTANT"

func main() {

	fmt.Println(s)
	const n = 500000000
	const d = 50500 / 2

	fmt.Println(d)
	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
