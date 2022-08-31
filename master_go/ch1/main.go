package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	for v:=1; v < len(os.Args); v++ {
		n, err := strconv.ParseFloat(os.Args[v], 64)
		if err != nil {
			continue
		}
		fmt.Println(n)
	}
}
