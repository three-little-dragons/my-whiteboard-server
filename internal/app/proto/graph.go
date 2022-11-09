package proto

type Paint[T Graph] struct {
	Version  string
	Type     string `json:"type"`
	Elements []T
}

type Graph struct {
	ID     int64 `json:"id"`
	UserId string
	X      float64
	Y      float64
}
