package geometry

type DCEL struct {
}

type Vertex struct {
	Coordinate   Point
	IncidentEdge *HalfEdge
}

type HalfEdge struct {
	Origin       *Vertex
	Twin         *HalfEdge
	IncidentFace *Face
	Next         *HalfEdge
	Prev         *HalfEdge
}

type Face struct {
	OuterComponent *HalfEdge
	InnerComponent []*HalfEdge
}
