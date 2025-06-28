package processor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractLocations(t *testing.T) {
	// Create temp directory
	dir := t.TempDir()

	// Sample GeoJSON Feature
	content := `
{
  "type": "Feature",
  "geometry": {
    "type": "Polygon",
    "coordinates": [
      [
        [1.0, 1.0],
        [3.0, 1.0],
        [3.0, 3.0],
        [1.0, 3.0],
        [1.0, 1.0]
      ]
    ]
  },
  "properties": {
    "adm3_en": "Testville"
  }
}`

	filePath := filepath.Join(dir, "test.json")
	if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write test file: %v", err)
	}

	// Run the function
	locations, err := ExtractLocations(dir)
	if err != nil {
		t.Fatalf("ExtractLocations returned error: %v", err)
	}

	// Assertions
	if len(locations) != 1 {
		t.Fatalf("expected 1 location, got %d", len(locations))
	}

	loc := locations[0]
	if loc.Municipality != "Testville" {
		t.Errorf("expected 'Testville', got '%s'", loc.Municipality)
	}

	if loc.Latitude != 2.0 || loc.Longitude != 2.0 {
		t.Errorf("expected centroid (2.0, 2.0), got (%f, %f)", loc.Latitude, loc.Longitude)
	}
}
