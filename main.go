package main

import (
	"fmt"
	"go_json_utility/guitar"
	"go_json_utility/parser"
	"go_json_utility/writer"
	"os"
)

func main() {

	// take command line arguments
	if (len(os.Args) < 3) {
		os.Exit(1)
	}

	switch (os.Args[1]) {
	case "read":
		var guitar guitar.Guitar

		err := parser.Parse(os.Args[2], &guitar)
	
		if (err != nil) {
			os.Exit(1)
		}
	
		fmt.Println("Name: ", guitar.Name)
		fmt.Println("Brand: ", guitar.Brand)
		fmt.Println("Year: ", guitar.Year)

	case "write":

		err := writer.Write(os.Args[2])

		if (err != nil) {
			os.Exit(1)
		}

	default:
	}


	os.Exit(0)
}