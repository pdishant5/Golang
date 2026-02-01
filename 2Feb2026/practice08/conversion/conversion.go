package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	floats := make([]float64, len(strings))

	for stringIndex, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("Error converting to float!")
		}

		floats[stringIndex] = floatPrice
	}
	return floats, nil
}
