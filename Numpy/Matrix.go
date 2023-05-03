package Numpy

import (
	"gonum.org/v1/gonum/mat"
	"log"
)

// CreateMatrixDenseFromArray 从数组创建矩阵/**
func CreateMatrixDenseFromArray(rowSize int, colSize int, arr []float64) *mat.Dense {
	mtx := mat.NewDense(rowSize, colSize, arr)
	log.Printf("\n%v", mat.Formatted(mtx))
	return mtx
}

// VisitDataByPosition 根据位置访问矩阵元素的值/**
func VisitDataByPosition(rowNo int, colNo int, mtx *mat.Dense) float64 {
	data := mtx.At(rowNo, colNo)
	log.Println(data)
	return data
}

// VisitRowByRowNo 根据行数访问矩阵元素的值/**
func VisitRowByRowNo(rowNo int, mtx *mat.Dense) []float64 {
	row := mat.Row(nil, rowNo, mtx)
	log.Println(row)
	return row
}

// VisitColByColNo 根据列数访问矩阵元素的值/**
func VisitColByColNo(colNo int, mtx *mat.Dense) []float64 {
	col := mat.Col(nil, colNo, mtx)
	log.Println(col)
	return col
}

// UpdateDataByPosition 根据位置修改矩阵元素的值/**
func UpdateDataByPosition(rowNo int, colNo int, mtx *mat.Dense, data float64) {
	log.Printf("Ori[%d][%d]: %v", rowNo, colNo, mtx.At(rowNo, colNo))
	mtx.Set(rowNo, colNo, data)
	log.Printf("Cur[%d][%d]: %v", rowNo, colNo, mtx.At(rowNo, colNo))
}

// UpdateRowByRowNo 根据行号修改矩阵元素的值/**
func UpdateRowByRowNo(rowNo int, mtx *mat.Dense, data []float64) {
	log.Printf("Ori[%d][-]: %v", rowNo, mat.Row(nil, rowNo, mtx))
	mtx.SetRow(rowNo, data)
	log.Printf("Cur[%d][-]: %v", rowNo, mat.Row(nil, rowNo, mtx))
}

// UpdateColByColNo 根据行号修改矩阵元素的值/**
func UpdateColByColNo(colNo int, mtx *mat.Dense, data []float64) {
	log.Printf("Ori[-][%d]: %v", colNo, mat.Col(nil, colNo, mtx))
	mtx.SetCol(colNo, data)
	log.Printf("Cur[-][%d]: %v", colNo, mat.Col(nil, colNo, mtx))
}
