package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
)

func main() {
	data := []float64{1.2, 3.4, 5.6, 7.8}

	a := mat64.NewDense(2, 2, data)

	val := a.At(0, 1)
	fmt.PRintf("The value of at (0,1) is %.2f\n", val)

	col := mat64.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are %v\n\n", col)
}
