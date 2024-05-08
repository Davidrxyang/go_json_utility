package main

import (
	"fmt"
	"os"
	"go_json_utility/guitar"
	"go_json_utility/parser"
	"go_json_utility/writer"

	"crypto/rand"

)

func main() {

	s := make([]byte, 10);
	rand.Read(s);
	fmt.Println("RUN ID: TEST-" + fmt.Sprintf("%x", s)[2 : 12]);

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