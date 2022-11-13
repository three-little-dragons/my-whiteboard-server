package proto

type FreeDraw struct {
	Graph
	Opacity     int
	FillStyle   string
	StrokeWidth float64
	StrokeColor float64
	UpdatedAt   int64
	Points      []Point
}
