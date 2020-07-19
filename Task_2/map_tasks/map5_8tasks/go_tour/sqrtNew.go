package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (err ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.2f", err)
}

func Sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNegativeSqrt(num)
	}
	z := 1.0
	z = z - (z*z-num)/(2*z)
	// допустимо любое значение большее чем 0.00001
	delta := 1.0
	for delta > 0.00001 {
		zPrev := z
		z = z - (z*z-num)/(2*z)
		delta = zPrev - z
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
