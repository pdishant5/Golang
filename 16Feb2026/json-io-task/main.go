package main

import (
	"flag"
	"log"

	"json-io/parser"
	"json-io/writer"
)

func main() {
	// CLI flags for optional configurations
	delimiter := flag.String("delimiter", ",", "Delimiter for CSV file")
	headers := flag.Bool("headers", true, "Include CSV headers")

	// Parse flags
	flag.Parse()

	// validating input and output files
	args := flag.Args()
	if len(args) < 2 {
		log.Println("Usage: JSONtoCSV <input.json> <output.csv> [options]")
		flag.PrintDefaults()
		return
	}

	// inline arguments for input and output files
	inputFile := args[0]
	outputFile := args[1]

	// fmt.Println("Input File:", inputFile)
	// fmt.Println("Output File:", outputFile)
	// fmt.Println("Delimiter:", *delimiter)
	// fmt.Println("Headers:", *headers)

	records, err := parser.ReadAndParseJSON(inputFile, outputFile, *delimiter, *headers)
	if err != nil {
		log.Println("Error reading the file:", err)
	}

	err = writer.WriteToFile(outputFile, *delimiter, *headers, records)
	if err != nil {
		log.Println("Error writing to file: " + err.Error())
	}
}
