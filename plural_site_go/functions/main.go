package main

import (
	"fmt"
)

func main() {

	type circle struct {
		r string
		d string
		c string
	}

	cir1 := circle{
		r: "Chad",
		d: "Jessica",
		c: "Dad",
	}
	fmt.Println(cir1.r)

}

