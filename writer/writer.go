package writer

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go_json_utility/guitar"
	"os"
	"strconv"
	"strings"
)

func Write(filename string) error {

	// prompt user to input from CLI

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter Name: ")
	name, err := reader.ReadString('\n')

	if (err != nil) {
		fmt.Println("Error input!")
		return err
	}

	name = trim(name)

	fmt.Println("Enter Brand: ")
	brand, err := reader.ReadString('\n')

	if (err != nil) {
		fmt.Println("Error input!")
		return err
	}

	brand = trim(brand)

	fmt.Println("Enter Year: ")
	yearString, err := reader.ReadString('\n')

	if (err != nil) {
		fmt.Println("Error input!")
		return err
	}

	yearString = trim(yearString)
	year, err := strconv.Atoi(yearString)

	if (err != nil) {
		fmt.Println("Error converting input to integer!")
		return err
	}

	guitar := guitar.Guitar{Name: name, Brand: brand, Year: year}

	jsonOutput, err := json.Marshal(guitar)

	if (err != nil) {
		fmt.Println("Error converting to json!")
		return err
	}

	file, err := os.Create(filename)
	
	if (err != nil) {
		fmt.Println("Error creating file!")
		return err
	}

	defer file.Close() // called after func execution?
	
	_, err = file.Write(jsonOutput)

	if (err != nil) {
		fmt.Println("Error writing to file!")
		return err
	}

	fmt.Printf("%s created\n", filename)

	return nil
}

func trim(str string) string{
	return strings.TrimSpace(str)
}