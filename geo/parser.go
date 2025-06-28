package geo

import "fmt"

func ConvertRawPolygon(data interface{}) (PolygonGeometry, error) {
	raw, ok := data.([]interface{})
	if !ok {
		return nil, fmt.Errorf("expected Polygon to be [][][2]float64")
	}

	var poly PolygonGeometry

	for _, ringRaw := range raw {
		ringIface, ok := ringRaw.([]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid ring in Polygon")
		}
		var ring LinearRing
		for _, ptRaw := range ringIface {
			pt, ok := ptRaw.([]interface{})
			if !ok || len(pt) < 2 {
				return nil, fmt.Errorf("invalid point in ring")
			}
			lon, ok1 := pt[0].(float64)
			lat, ok2 := pt[1].(float64)
			if !ok1 || !ok2 {
				return nil, fmt.Errorf("point coords must be float64")
			}
			ring = append(ring, Point{lon, lat})
		}
		poly = append(poly, ring)
	}

	return poly, nil
}
