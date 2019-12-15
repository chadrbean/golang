package main

import (
    "encoding/csv"
    "fmt"
	"os"
	"log"
	"io"
	"strings"
)



func main() {
	csvFile, _ := os.Open("file.csv") // the underscore is undefined placeholder for error
	reader := csv.NewReader(csvFile) // read csv after reading it above
	// fmt.Print(reader)
	for {
		line, error := reader.Read() // read everyline line by lin, if error its at end of file
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		for i:=0; i<=len(line)-1; i++ {
			s := strings.SplitN(line, "+", 2)
			fmt.Print(s)
		}
		// fmt.Print(line[0:], "\n")
		// s := strings.Split(line, "+")
		// fmt.Print(s)
	}
}