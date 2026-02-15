package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func ReadAndParseJSON(inputFile, outputFile, delimiter string, headers bool) ([]map[string]interface{}, error) {
	data, err := os.ReadFile(inputFile)
	// fmt.Println("Data:", string(data))
	if err != nil {
		// log.Fatalln("Error reading the file:", err)
		return nil, errors.New("Error reading the file: " + err.Error())
	}

	// de-serializing the JSON data
	var raw interface{}
	err = json.Unmarshal(data, &raw)
	// fmt.Println("Raw data:", raw)
	if err != nil {
		// log.Fatalln("Error parsing the file:", err)
		return nil, errors.New("Error parsing the file: " + err.Error())
	}

	switch v := raw.(type) {
	case []interface{}:
		fmt.Println("JSON is an array of objects!")
		records := make([]map[string]interface{}, 0)

		for i, item := range v {
			obj, ok := item.(map[string]interface{})
			if !ok {
				// log.Fatalf("Array item %d is not an object", i)
				return nil, errors.New(fmt.Errorf("Array item %d is not an object", i).Error())
			}
			records = append(records, obj)
		}

		fmt.Printf("Parsed %d records!\n", len(records))
		// fmt.Println("Records:", records)
		return records, nil

	case map[string]interface{}:
		fmt.Println("JSON is a single object!")
		records := []map[string]interface{}{v}

		fmt.Printf("Parsed %d record!\n", len(records))
		// fmt.Println("Records:", records)
		return records, nil

	default:
		return nil, errors.New("Unsupported JSON format!")
	}
}
