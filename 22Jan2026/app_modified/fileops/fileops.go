package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return defaultValue, errors.New("Failed to find the file!")
	}

	floatText := string(data)
	floatValue, err := strconv.ParseFloat(floatText, 64)

	if err != nil {
		return defaultValue, errors.New("Failed to parse the value!")
	}
	return floatValue, nil
}

func WriteFloatToFile(fileName string, floatValue float64) {
	floatText := fmt.Sprint(floatValue)
	os.WriteFile(fileName, []byte(floatText), 0644)
}

// capitalized names are exported
// var Name = "fileops_package" // accessed across packages
// var name = "fileops_package" // accessed within the package
