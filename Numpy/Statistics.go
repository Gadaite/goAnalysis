package Numpy

import (
	"gonum.org/v1/gonum/stat"
	"log"
)

// ChiSquareEachArray 计算两个数组的卡方值/**
func ChiSquareEachArray(a1 []float64, a2 []float64) float64 {
	chiSquare := stat.ChiSquare(a1, a2)
	log.Printf("chiSquare: %f", chiSquare)
	return chiSquare
}

func P() {

}
