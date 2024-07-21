package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
)

func main() {
	data := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	e := mat64.NewDense(3, 3, data)
	fa := mat64.Formatted(e, mat64.Prefix("    "))
	fmt.Printf("A = %v\n\n", fa)
}
