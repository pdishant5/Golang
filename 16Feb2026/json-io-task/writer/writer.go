package writer

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
	We can write separate functions for each delimiter type,or
	we can also have a single function that manually writes to a file using the specified delimiter.
	For simplicity, i'm implementing the first approach.
	I will extend it to support more delimiters in the future without duplicating code.
*/

func WriteToFile(outputFile, delimiter string, includeHeaders bool, records []map[string]interface{}) error {
	switch delimiter {
	case ",":
		return writeToCSV(outputFile, includeHeaders, records)
	case "\t":
		return writeToTSV(outputFile, includeHeaders, records)
	case "|":
		return writeToPipeSeparated(outputFile, includeHeaders, records)
	default:
		return errors.New("Unsupported delimiter: " + delimiter)
	}
}

func writeToCSV(outputFile string, includeHeaders bool, records []map[string]interface{}) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := collectHeaders(records)
	if includeHeaders {
		csvWriter.Write(headers)
		// fmt.Println("Headers:", headers)
	}

	for _, record := range records {
		row := make([]string, len(headers))
		for i, header := range headers {
			row[i] = valueToString(record[header])
		}
		csvWriter.Write(row)
		// fmt.Println("Row:", row)
	}
	fmt.Printf("Data successfully written to the file: %s!\n", outputFile)
	return nil
}

func collectHeaders(records []map[string]interface{}) []string {
	set := make(map[string]struct{})

	for _, record := range records {
		for key := range record {
			set[key] = struct{}{}
		}
	}

	headers := make([]string, 0, len(set))
	for key := range set {
		headers = append(headers, key)
	}
	sort.Strings(headers)
	return headers
}

func valueToString(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case string:
		return val
	case int:
		return strconv.FormatInt(int64(val), 10)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

func writeToTSV(outputFile string, includeHeaders bool, records []map[string]interface{}) error {
	// similar to WriteToCSV but with tab delimiter
	return nil
}

func writeToPipeSeparated(outputFile string, includeHeaders bool, records []map[string]interface{}) error {
	// similar to WriteToCSV but with pipe delimiter
	return nil
}
