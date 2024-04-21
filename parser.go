package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Guitar struct {
	Name string `json:"name`
	Brand string `json:brand`
	Year int `json:year`
}

func parse(filename string) error {
	
	data, err := os.ReadFile(filename)

	if (err != nil) {
		fmt.Println("Error reading file!")
		return err
	}

	var guitar Guitar

	err = json.Unmarshal(data, &guitar)

	if (err != nil) {
		fmt.Println("Error parsing json!")
		return err
	}

	fmt.Println("Name: ", guitar.Name)
	fmt.Println("Brand: ", guitar.Brand)
	fmt.Println("Year: ", guitar.Year)

	return nil
}

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