package Numpy

import (
	"gonum.org/v1/gonum/mat"
	"math"
	"reflect"
	"testing"
)

func TestCreateMatrixDenseFromArray(t *testing.T) {
	type args struct {
		rowSize int
		colSize int
		arr     []float64
	}
	tests := []struct {
		name string
		args args
		want *mat.Dense
	}{
		// TODO: Add test cases.
		{name: "case 1", args: args{rowSize: 2, colSize: 1, arr: []float64{1, 3}}, want: mat.NewDense(2, 1, []float64{1, 3})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMatrixDenseFromArray(tt.args.rowSize, tt.args.colSize, tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateMatrixDenseFromArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisitDataByPosition(t *testing.T) {
	type args struct {
		rowNo int
		colNo int
		mtx   *mat.Dense
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"case 1", args{1, 0, mat.NewDense(2, 1, []float64{1, 2})}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VisitDataByPosition(tt.args.rowNo, tt.args.colNo, tt.args.mtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VisitDataByPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateColByColNo(t *testing.T) {
	type args struct {
		colNo int
		mtx   *mat.Dense
		data  []float64
	}
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "case 1", args: args{0, m, []float64{math.NaN(), 100, 100}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateColByColNo(tt.args.colNo, tt.args.mtx, tt.args.data)
		})
	}
}

func TestUpdateDataByPosition(t *testing.T) {
	type args struct {
		rowNo int
		colNo int
		mtx   *mat.Dense
		data  float64
	}
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case 1", args{1, 1, m, 55}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateDataByPosition(tt.args.rowNo, tt.args.colNo, tt.args.mtx, tt.args.data)
		})
	}
}

func TestUpdateRowByRowNo(t *testing.T) {
	type args struct {
		rowNo int
		mtx   *mat.Dense
		data  []float64
	}
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case 1", args{0, m, []float64{11, 22, 33}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateRowByRowNo(tt.args.rowNo, tt.args.mtx, tt.args.data)
		})
	}
}

func TestVisitColByColNo(t *testing.T) {
	type args struct {
		colNo int
		mtx   *mat.Dense
	}
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{"case 1", args{1, m}, []float64{2, 5, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VisitColByColNo(tt.args.colNo, tt.args.mtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VisitColByColNo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVisitRowByRowNo(t *testing.T) {
	type args struct {
		rowNo int
		mtx   *mat.Dense
	}
	m := mat.NewDense(3, 3, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9})
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{"case 1", args{1, m}, []float64{4, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VisitRowByRowNo(tt.args.rowNo, tt.args.mtx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VisitRowByRowNo() = %v, want %v", got, tt.want)
			}
		})
	}
}
