package Numpy

import (
	"fmt"
	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
	"log"
)

// DotArrays 对两个数组计算内积/**
func DotArrays(vectorA []float64, vectorB []float64) float64 {
	dot := floats.Dot(vectorA, vectorB)
	return dot
}

// CreateVecDenseFromArray 从数组创建向量/**
func CreateVecDenseFromArray(array []float64) *mat.VecDense {
	dense := mat.NewVecDense(len(array), array)
	fmt.Println(dense)
	return dense
}

// ScaleVecByValue 缩放向量/**
func ScaleVecByValue(vec *mat.VecDense, scale float64, useOri bool) *mat.VecDense {
	tmp := mat.NewVecDense(vec.Len(), nil)
	if useOri {
		tmp = vec
	} else {
		tmp.CopyVec(vec)
	}
	tmp.ScaleVec(scale, tmp)
	if &tmp != &vec {
		log.Println("新建向量并转换")
	} else {
		log.Println("对原向量直接转换")
	}
	return tmp
}

// Norm2VecByValue 计算向量2-范数/**
func Norm2VecByValue(vec *mat.VecDense) float64 {
	nrm2 := blas64.Nrm2(vec.RawVector())
	log.Println(nrm2)
	return nrm2
}
