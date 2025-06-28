package processor

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"latlon-generator/geo"
)

func ExtractLocations(inputDir string) ([]Location, error) {
	var results []Location

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		fileData, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		var raw map[string]interface{}
		if err := json.Unmarshal(fileData, &raw); err != nil {
			log.Printf("Failed to parse JSON in %s: %v", path, err)
			return nil
		}

		featuresRaw := []interface{}{}
		if raw["type"] == "FeatureCollection" {
			if fList, ok := raw["features"].([]interface{}); ok {
				featuresRaw = fList
			}
		} else if raw["type"] == "Feature" {
			featuresRaw = append(featuresRaw, raw)
		} else {
			log.Printf("Skipping file: unknown GeoJSON type: %v", raw["type"])
			return nil
		}

		for _, f := range featuresRaw {
			feature := f.(map[string]interface{})
			props := feature["properties"].(map[string]interface{})
			geometry := feature["geometry"].(map[string]interface{})

			if geometry["type"] != "Polygon" || geometry["coordinates"] == nil {
				continue
			}

			polygon, err := geo.ConvertRawPolygon(geometry["coordinates"])
			if err != nil || len(polygon) == 0 {
				continue
			}
			centroid := geo.CalculatePolygonCentroid(polygon[0])

			results = append(results, Location{
				Province:     "", // Optionally read from props["adm1_en"]
				Municipality: props["adm3_en"].(string),
				Latitude:     centroid[1],
				Longitude:    centroid[0],
			})
		}
		return nil
	})

	return results, err
}
