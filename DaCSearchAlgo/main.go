package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var inputs [][]int

// takes in search value and array to search
func dAC(searchNum int, in []int) bool {
	if len(in) <= 10 {
		// search without dividing if less than 10 numbers
		for _, num := range in {
			if num == searchNum {
				return true
			}
		}
		return false
	}

	// divide array in to 3 equal parts
	var one []int
	var two []int
	var three []int
	for i, num := range in {
		if i < len(in)/3 {
			one = append(one, num)
		} else if i < len(in)*2/3 {
			two = append(two, num)
		} else {
			three = append(three, num)
		}
	}

	// call dAC on each sub division and return
	return dAC(searchNum, one) || dAC(searchNum, two) || dAC(searchNum, three)
}

func main() {
	// check for valid arg number
	if len(os.Args[1:]) != 1 {
		log.Fatalf("Expected 1 file path arg, got %d args", len(os.Args[1:]))
	}

	// read file per line
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// 64k line length maximum
	var strArry []string
	var intArry = []int{}
	// parsae file in to arrays
	for scanner.Scan() {
		strArry = strings.Split(scanner.Text(), ",")
		for _, i := range strArry {
			num, err := strconv.Atoi(i)
			if err != nil {
				log.Fatal(err)
			}
			intArry = append(intArry, num)
		}
		inputs = append(inputs, intArry)
		intArry = []int{}
	}

	// call dac and return
	for _, numArray := range inputs {
		result := ""
		searchNum := numArray[0]
		if !dAC(searchNum, numArray[1:]) {
			result = " not"
		}

		fmt.Printf("The number %v was%v found in the array %v\n", searchNum, result, numArray[1:])
	}
}
