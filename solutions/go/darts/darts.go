package darts

import "math"

func IsInCircle(x, y, radius float64) bool {
    return math.Sqrt((x * x) + (y * y)) <= radius
}

func Score(x, y float64) int {
    if (IsInCircle(x, y, 1)) {
        return 10
    }
    if (IsInCircle(x, y, 5)) {
        return 5
    }
    if (IsInCircle(x, y, 10)) {
        return 1
    }
    return 0
}
