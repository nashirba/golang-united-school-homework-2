package square

import "math"

var (
	SidesTriangle int = 3
	SidesSquare int = 4
	SidesCircle int = 0
)
type CustomInt int

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	var Square float64
	switch int(sidesNum) {
	case SidesCircle:
		Square = math.Pi *sideLen * sideLen
	case SidesTriangle:
		Square = math.Pow(0.75, 0.5) * sideLen * sideLen
	case SidesSquare:
		Square = sideLen * sideLen
	}
	return Square
}
