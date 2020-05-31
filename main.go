package main

import (
	"fmt"

	"github.com/homework2/code"
	"github.com/homework2/pig"
)

func main() {
	fmt.Println(pig.PigLatin("banana"))
	fmt.Println(code.Encode("hello"))
	fmt.Println(code.Decode("h3 th2r2"))
}
