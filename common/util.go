package common

import (
	"math"
	"strconv"
)

func ParseInt(str string) (int64, error) {
	num, err := strconv.ParseInt(str, 10, 64)
	return num, err
}

func Round(s float64) int64 {
	return int64(math.Ceil(s))
}
