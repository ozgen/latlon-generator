package processor

import (
	"strings"
	"testing"
)

func TestGenerateSQL(t *testing.T) {
	locations := []Location{
		{
			Province:     "Test's Province",
			Municipality: "O'Neil Town",
			Latitude:     10.1234,
			Longitude:    123.4567,
		},
	}

	table := "locations"
	sql := GenerateSQL(locations, table)

	// Check CREATE TABLE
	if !strings.Contains(sql, "CREATE TABLE IF NOT EXISTS locations") {
		t.Error("missing CREATE TABLE statement")
	}

	// Check INSERT statement
	if !strings.Contains(sql, "INSERT INTO locations") {
		t.Error("missing INSERT statement")
	}

	// Check escape
	if !strings.Contains(sql, "Test''s Province") || !strings.Contains(sql, "O''Neil Town") {
		t.Error("expected SQL to escape single quotes")
	}

	// Check coordinates
	if !strings.Contains(sql, "10.123400") || !strings.Contains(sql, "123.456700") {
		t.Error("expected coordinates in output")
	}
}
