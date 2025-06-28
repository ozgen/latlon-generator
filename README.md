# LatLon Generator

[![Test Coverage](https://github.com/ozgen/latlon-generator/actions/workflows/ci.yml/badge.svg)](https://github.com/ozgen/latlon-generator/actions/workflows/ci.yml)
[![Latest Release](https://img.shields.io/github/v/release/ozgen/latlon-generator?label=latest)](https://github.com/ozgen/latlon-generator/releases/latest)

A Go command-line tool that parses GeoJSON files and computes geographic centroids (latitude and longitude). It can output either:

- A JSON array of location records
- SQL statements to create and populate a database table

---

## Features

- Reads single or FeatureCollection GeoJSON files
- Computes centroids from raw polygon coordinates
- Outputs either JSON or SQL
- CLI-friendly with flag-based configuration

---

## Usage

### Build

```bash
go build -o latlon
```

### Run (default: JSON)

```bash
./bin/latlon -input=/path/to/geojson -output=locations.json
```

### Run with SQL output

```bash
./bin/latlon -input=/path/to/geojson -output=locations.sql --sql --table=locations
```

---

## Testing

```bash
make test
```

To view line-by-line test coverage in browser:

```bash
go test ./... -coverprofile=coverage.out
go tool cover -html=coverage.out
```

---

## Output Format

### JSON

```json
[
  {
    "province": "",
    "municipality": "Adams",
    "latitude": 18.5284,
    "longitude": 120.9018
  }
]
```

### SQL

```sql
CREATE TABLE IF NOT EXISTS lookup_location (
  id SERIAL PRIMARY KEY,
  province TEXT,
  municipality TEXT NOT NULL,
  latitude DOUBLE PRECISION NOT NULL,
  longitude DOUBLE PRECISION NOT NULL
);

INSERT INTO lookup_location (province, municipality, latitude, longitude)
VALUES ('', 'Adams', 18.5284, 120.9018);
```

---

## Structure

```
latlon-generator/
├── main.go
├── geo/                # Centroid and polygon parsers
├── processor/          # File walker, SQL and JSON emitters
├── output.json/sql     # Generated results
```

---

