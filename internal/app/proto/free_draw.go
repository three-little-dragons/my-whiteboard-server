package proto

type Point struct {
	X float64
	Y float64
}

type FreeDraw struct {
	Graph
	Opacity     int
	FillStyle   string
	StrokeWidth float64
	StrokeColor float64
	UpdatedAt   int64
	Points      []Point
}
