package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func min[T Number](vals ...T) T {
	if len(vals) == 0 {
		panic("min function accept non-zero number of arguments")
	}

	min_val := vals[0]

	for _, v := range vals {
		if min_val > v {
			min_val = v
		}
	}

	return min_val
}

func max[T Number](vals ...T) T {
	if len(vals) == 0 {
		panic("max function accept non-zero number of arguments")
	}

	max_val := vals[0]

	for _, v := range vals {
		if v > max_val {
			max_val = v
		}
	}

	return max_val
}

func Constraint[T Number](val, left, right T) T {
	return max(min(val, right), left)
}
