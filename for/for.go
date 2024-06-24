package main

import "fmt"

func main() {
	i := 1

	fmt.Println(">>>")
	for i <= 3 {
		fmt.Print(i)
		i = i + 1
	}

	fmt.Println("\n>>>")

	for j := 0; j < 3; j++ {
		fmt.Print(j)
	}
	fmt.Println("\n>>>")
	for {
		fmt.Print("loop")
		break

	}
	fmt.Println("\n>>>")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Print(n)
	}
}
