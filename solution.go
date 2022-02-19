package square

import "math"

type CustomInt int
var (
	SidesTriangle CustomInt = 3
	SidesSquare CustomInt = 4
	SidesCircle CustomInt = 0
)

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	var Square float64
	switch sidesNum {
	case SidesCircle:
		Square = math.Pi *sideLen * sideLen
	case SidesTriangle:
		Square = math.Pow(0.75, 0.5) * sideLen * sideLen
	case SidesSquare:
		Square = sideLen * sideLen
	}
	return Square
}
