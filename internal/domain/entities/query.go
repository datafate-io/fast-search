package entities

type Query struct {
	Text     string
	Vector   []float64
	ID       string
	VectorID string
}

func (q *Query) SetVector(vector []float64) {
	q.Vector = vector
}

func (q *Query) GetVector() []float64 {
	return q.Vector
}

func (q *Query) IsValidQuery() bool {
	return len(q.Text) > 0 || len(q.Vector) < 256
}
