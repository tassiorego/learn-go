package main

import (
	"fmt"
	"log"

	"github.com/ortense/consolestyle"
)

func main() {
	fmt.Println(consolestyle.Green("Hello, World!"))
	log.Println("Hello, World!")
}
