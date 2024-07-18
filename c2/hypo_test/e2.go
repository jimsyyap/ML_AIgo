package main

import (
	"fmt"
	"github.com/gonum/stat"
	"github.com/gonum/stat/distuv"
)

func main() {
	observed := []float64{
		260.0,
		135.0,
		105.0,
	}

	totalObserved := []float64{
		totalObserved * 0.60,
		totalObserved * 0.25,
		totalObserved * 0.15,
	}

	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Printf("\nChi-square: %0.2f\n", chiSquare)

	chiDist := distuv.ChiSquared{
		K:   2.0,
		Src: nil,
	}

	pValue := chiDist.Prob(chiSquare)
	fmt.Printf("\nP-value: %0.4f\n", pValue)
}
