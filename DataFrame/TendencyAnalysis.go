package DataFrame

import (
	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
	"log"
	"strconv"
)

// CentralTendencyStruct 集中趋势统计结果/**
type CentralTendencyStruct struct {
	ColName    string
	Mean       float64
	Mode       float64
	ModelCount float64
	Median     float64
}

// DiscreteTendencyStruct 离中趋势统计结果/**
type DiscreteTendencyStruct struct {
	ColName    string
	Max        float64
	Min        float64
	Range      float64
	Variance   float64
	Std        float64
	Quantile25 float64
	Quantile50 float64
	Quantile75 float64
}

// CentralTendency 集中趋势统计/**
func CentralTendency(df dataframe.DataFrame, colName string, decimalPlaces int) CentralTendencyStruct {
	colDf := df.Col(colName)
	colData := colDf.Float()
	mean := stat.Mean(colData, nil)
	mode, count := stat.Mode(colData, nil)
	median, err := stats.Median(colData)
	if err != nil {
		log.Fatal(err)
	}
	var res CentralTendencyStruct
	if decimalPlaces > 0 {
		res = CentralTendencyStruct{
			ColName:    colName,
			Mean:       decimalFloat64(decimalPlaces, mean),
			Mode:       decimalFloat64(decimalPlaces, mode),
			ModelCount: decimalFloat64(decimalPlaces, count),
			Median:     decimalFloat64(decimalPlaces, median),
		}
	} else {
		res = CentralTendencyStruct{
			ColName:    colName,
			Mean:       mean,
			Mode:       mode,
			ModelCount: count,
			Median:     median,
		}
	}
	log.Printf("\n%v", res)
	return res
}

// DiscreteTendency 离中趋势统计/**
func DiscreteTendency(df dataframe.DataFrame, colName string, decimalPlaces int) DiscreteTendencyStruct {
	colDf := df.Col(colName)
	colData := colDf.Float()
	max := floats.Max(colData)
	min := floats.Min(colData)
	rangeVal := max - min
	variance, err := stats.Variance(colData)
	if err != nil {
		log.Fatal(err)
	}
	stdDev := stat.StdDev(colData, nil)
	floats.Argsort(colData, make([]int, len(colData)))
	quantile25 := stat.Quantile(0.25, stat.Empirical, colData, nil)
	quantile50 := stat.Quantile(0.50, stat.Empirical, colData, nil)
	quantile75 := stat.Quantile(0.75, stat.Empirical, colData, nil)
	var res DiscreteTendencyStruct
	if decimalPlaces <= 0 {
		res = DiscreteTendencyStruct{
			ColName:    colName,
			Max:        max,
			Min:        min,
			Range:      rangeVal,
			Variance:   variance,
			Std:        stdDev,
			Quantile25: quantile25,
			Quantile50: quantile50,
			Quantile75: quantile75,
		}
	} else {
		res = DiscreteTendencyStruct{
			ColName:    colName,
			Max:        decimalFloat64(decimalPlaces, max),
			Min:        decimalFloat64(decimalPlaces, min),
			Range:      decimalFloat64(decimalPlaces, rangeVal),
			Variance:   decimalFloat64(decimalPlaces, variance),
			Std:        decimalFloat64(decimalPlaces, stdDev),
			Quantile25: decimalFloat64(decimalPlaces, quantile25),
			Quantile50: decimalFloat64(decimalPlaces, quantile50),
			Quantile75: decimalFloat64(decimalPlaces, quantile75),
		}
	}

	log.Printf("\n%v", res)
	return res
}

// decimalFloat64 保留指定位数小数/**
func decimalFloat64(decimalPlaces int, data float64) float64 {
	res, err := strconv.ParseFloat(strconv.FormatFloat(data, 'f', decimalPlaces, 64), 64)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
