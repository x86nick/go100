package datafile

import (
	"bufio"
	"os"
	"strconv"
)

// GetFloats reads a float64 from each line from a file.
// function returns an array of float64 value
//func GetFloats(fileName string) ([3]float64, error) { // switch to return a slice
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64 // contains nil by default
	file, err := os.Open(fileName)
	if err != nil {
		return numbers, err
	}
	i := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		number, err := strconv.ParseFloat(scanner.Text(), 64) // convert string to float64 and assing it to temp variable
		if err != nil {
			return numbers, err
		}
		// i++
		numbers = append(numbers, number) // append the new number to the slice
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {

		return numbers, err
	}
	return numbers, nil
}
