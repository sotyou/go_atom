package atom

import "fmt"

func S(v fmt.Stringer) string {
	return v.String()
}

type Integer interface {
	Int() int
}

func I(i Integer) int {
	return i.Int()
}

type Floater interface {
	Float() float64
}

func F(i Floater) float64 {
	return i.Float()
}
