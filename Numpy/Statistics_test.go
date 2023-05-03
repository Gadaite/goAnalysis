package Numpy

import (
	"gonum.org/v1/gonum/stat"
	"testing"
)

func TestChiSquareEachArray(t *testing.T) {
	type args struct {
		a1 []float64
		a2 []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"case 1", args{[]float64{1, 2}, []float64{3, 4}}, stat.ChiSquare([]float64{1, 2}, []float64{3, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChiSquareEachArray(tt.args.a1, tt.args.a2); got != tt.want {
				t.Errorf("ChiSquareEachArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
