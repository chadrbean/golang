package main

import (
	"fmt"
)
func main(){
	sortArray := [...]int{5,10,6,3,7,8, 10, 24, 44, 54, 2, 4, 56}
	findLen := len(sortArray)
	preSort := [2]int{} //create array size 2 for sorting each pair
	for ii:=0; ii <= 6; ii++ {
		fmt.Print(ii)
		for i := 0; i <= findLen; i++ {
			if i == findLen -1 {
				break
			}
			if sortArray[i] > sortArray[i+1] {
				preSort[0] = sortArray[i+1]
				preSort[1] = sortArray[i]
				sortArray[i] = preSort[0]
				sortArray[i+1] = preSort[1]
			}
		}
		fmt.Print(sortArray, "\n")
	}
	
}