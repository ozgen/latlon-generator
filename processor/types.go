package processor

type Location struct {
	Province     string  `json:"province"`
	Municipality string  `json:"municipality"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}
