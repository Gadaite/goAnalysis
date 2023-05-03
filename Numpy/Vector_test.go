package Numpy

import (
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/mat"
	"reflect"
	"testing"
)

func TestDotArrays(t *testing.T) {
	type args struct {
		vectorA []float64
		vectorB []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"case 1", args{vectorA: []float64{1, 2, 3}, vectorB: []float64{1, 2, 3}}, 14},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DotArrays(tt.args.vectorA, tt.args.vectorB); got != tt.want {
				t.Errorf("DotArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateVecDenseFromArray(t *testing.T) {
	type args struct {
		array []float64
	}
	tests := []struct {
		name string
		args args
		want *mat.VecDense
	}{
		// TODO: Add test cases.
		{name: "case 1", args: args{array: []float64{1, 2, 3}}, want: mat.NewVecDense(3, []float64{1, 2, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateVecDenseFromArray(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateVecDenseFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScaleVecByValue(t *testing.T) {
	type args struct {
		vec    *mat.VecDense
		scale  float64
		useOri bool
	}
	tests := []struct {
		name string
		args args
		want *mat.VecDense
	}{
		// TODO: Add test cases.
		{name: "case 1", args: args{mat.NewVecDense(2, []float64{1, 2}), 2, true}, want: mat.NewVecDense(2, []float64{2, 4})},
		{name: "case 2", args: args{mat.NewVecDense(2, []float64{1, 2}), 2, false}, want: mat.NewVecDense(2, []float64{2, 4})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScaleVecByValue(tt.args.vec, tt.args.scale, tt.args.useOri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScaleVecByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNorm2VecByValue(t *testing.T) {
	type args struct {
		vec *mat.VecDense
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{name: "case 1", args: args{vec: mat.NewVecDense(2, []float64{1, 2})}, want: blas64.Nrm2(mat.NewVecDense(2, []float64{1, 2}).RawVector())},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Norm2VecByValue(tt.args.vec); got != tt.want {
				t.Errorf("Norm2VecByValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
