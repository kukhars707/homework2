package main

import (
	"fmt"

	"github.com/homework2/code"
	"github.com/homework2/pig"
)

func main() {
	pig, err := pig.PigLatin("banana")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(pig)
	fmt.Println(code.Encode("hello"))
	fmt.Println(code.Decode("h3 th2r2"))
}
