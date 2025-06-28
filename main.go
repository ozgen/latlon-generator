package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"latlon-generator/processor"
	"log"
	"os"
)

func main() {
	inputDir := flag.String("input", "/path/to/geojson", "Directory containing GeoJSON files")
	outputFile := flag.String("output", "output_latlon.json", "Output file name")
	outputSQL := flag.Bool("sql", false, "Output SQL instead of JSON")
	tableName := flag.String("table", "lookup_location", "SQL table name (used only if -sql is true)")
	flag.Parse()

	locations, err := processor.ExtractLocations(*inputDir)
	if err != nil {
		log.Fatalf("processing failed: %v", err)
	}

	if *outputSQL {
		sqlOut := processor.GenerateSQL(locations, *tableName)
		if err := os.WriteFile(*outputFile, []byte(sqlOut), 0644); err != nil {
			log.Fatalf("failed to write SQL file: %v", err)
		}
		fmt.Printf("SQL written to %s\n", *outputFile)
	} else {
		jsonOut, err := json.MarshalIndent(locations, "", "  ")
		if err != nil {
			log.Fatalf("failed to marshal JSON: %v", err)
		}
		if err := os.WriteFile(*outputFile, jsonOut, 0644); err != nil {
			log.Fatalf("failed to write JSON file: %v", err)
		}
		fmt.Printf("JSON written to %s\n", *outputFile)
	}
}
