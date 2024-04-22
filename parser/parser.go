package parser

import (
	"encoding/json"
	"fmt"
	"os"
	"go_json_utility/guitar"
)

func Parse(filename string, guitar* guitar.Guitar) error {
	
	data, err := os.ReadFile(filename)

	if (err != nil) {
		fmt.Println("Error reading file!")
		return err
	}

	err = json.Unmarshal(data, &guitar)

	if (err != nil) {
		fmt.Println("Error parsing json!")
		return err
	}



	return nil
}

/*

func main() {

	// take command line arguments
	if (len(os.Args) < 2) {
		os.Exit(1)
	}

	err := parse(os.Args[1])

	if (err != nil) {
		os.Exit(1)
	}

	os.Exit(0)
}

*/