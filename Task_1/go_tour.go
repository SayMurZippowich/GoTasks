package main

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

//Task: loops and functions
func Sqrt(num float64) float64 {
	z := 1.0
	z = z - (z*z-num)/(2*z)
	// допустимо любое значение большее чем 0.00001
	delta := 1.0
	for delta > 0.00001 {
		zPrev := z
		z = z - (z*z-num)/(2*z)
		delta = zPrev - z
	}
	return z
}

//Task: slices
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		tmpRow := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			tmpRow[x] = uint8((x + y) / 2)
		}
		pic[y] = tmpRow
	}
	return pic
}

//Task: map
func WordCount(s string) map[string]int {

	strArr := strings.Fields(s)
	wMap := make(map[string]int)

	for _, val := range strArr {
		wMap[val]++
	}
	return wMap
}

//Task: Fibonacci
func Fibonacci() func() int {
	prevVal := -1
	val := 1
	return func() int {
		buf := val
		val += prevVal
		prevVal = buf
		return val
	}
}

func main() {
	//Task: loops and functions
	fmt.Println("|math.Sqrt result:")
	fmt.Println(math.Sqrt(2))
	fmt.Println("|diy Sqrt result:")
	fmt.Println(Sqrt(2))
	//Task: slices
	fmt.Println("|Picture:")
	pic.Show(Pic)
	//Task: map
	fmt.Println("|Testing map task:")
	wc.Test(WordCount)
	//Task: Fibonacci
	fmt.Println("|Fibonacci number:")
	f := Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
