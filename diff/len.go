package diff

// Len enables a nil length so we can set defaults on nil
type Len struct {
	Value int
}

func NewLen(n int) *Len {
	return &Len{Value: n}
}
