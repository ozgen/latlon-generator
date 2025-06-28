package geo

type Point [2]float64

type LinearRing []Point

type PolygonGeometry []LinearRing

type MultiPolygonGeometry []PolygonGeometry
