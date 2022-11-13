package proto

// Proto is basic data transfer structure
type Proto struct {
	Type string
}

const (
	TypePaint  = "paint"
	TypeCursor = "cursor"
	TypeGraph  = "graph"
)

type Point struct {
	X float64
	Y float64
}

// Cursor represents touch/cursor position
type Cursor struct {
	Proto
	Point
}

type Paint struct {
	Proto
	Version string
	DrawTab DrawTab
	// userId -> []Graph
	Elements map[int64][]Graph
}

type Graph struct {
	ID int64
	Point
}
