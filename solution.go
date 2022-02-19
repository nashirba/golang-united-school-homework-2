package square

import "math"

type CustomInt int

func CalcSquare(sideLen float64, sidesNum CustomInt) float64 {
	var Square float64
	switch sidesNum {
	case 0:
		Square = math.Pi *sideLen * sideLen
	case 3:
		Square = math.Pow(0.75, 0.5) * sideLen * sideLen
	case 4:
		Square = sideLen * sideLen
	}
	return Square
}
