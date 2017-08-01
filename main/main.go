package main

import (
	"fmt"

	"github.com/williamsharkey/enterprise"
)

func main() {
	fmt.Println("hello from main")
	enterprise.Add(1, 2)
}

//Answer to life is the number 42.
func Answer() int {
	return 42
}
