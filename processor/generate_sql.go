package processor

import (
	"fmt"
	"strings"
)

func escape(s string) string {
	return strings.ReplaceAll(s, "'", "''")
}

func GenerateSQL(locations []Location, table string) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (\n", table))
	sb.WriteString("  id SERIAL PRIMARY KEY,\n")
	sb.WriteString("  province TEXT,\n")
	sb.WriteString("  municipality TEXT NOT NULL,\n")
	sb.WriteString("  latitude DOUBLE PRECISION NOT NULL,\n")
	sb.WriteString("  longitude DOUBLE PRECISION NOT NULL\n")
	sb.WriteString(");\n\n")

	for _, loc := range locations {
		stmt := fmt.Sprintf(
			"INSERT INTO %s (province, municipality, latitude, longitude) VALUES ('%s', '%s', %f, %f);\n",
			table,
			escape(loc.Province),
			escape(loc.Municipality),
			loc.Latitude,
			loc.Longitude,
		)
		sb.WriteString(stmt)
	}

	return sb.String()
}
