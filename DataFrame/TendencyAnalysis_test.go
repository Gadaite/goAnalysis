package DataFrame

import (
	"github.com/go-gota/gota/dataframe"
	"goAnalysisi/DataSource"
	"reflect"
	"testing"
)

func TestCentralTendency(t *testing.T) {
	data := DataSource.ReadCSVDataFrame("../Resource/DataSet/iris-data.csv")
	type args struct {
		df            dataframe.DataFrame
		colName       string
		decimalPlaces int
	}
	tests := []struct {
		name string
		args args
		want CentralTendencyStruct
	}{
		// TODO: Add test cases.
		{"case 1", args{data, "petal-length", 2}, CentralTendencyStruct{
			ColName:    "petal-length",
			Mean:       3.76,
			Mode:       1.50,
			ModelCount: 14.00,
			Median:     4.35,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CentralTendency(tt.args.df, tt.args.colName, tt.args.decimalPlaces); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CentralTendency() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiscreteTendency(t *testing.T) {
	data := DataSource.ReadCSVDataFrame("../Resource/DataSet/iris-data.csv")
	type args struct {
		df            dataframe.DataFrame
		colName       string
		decimalPlaces int
	}
	tests := []struct {
		name string
		args args
		want DiscreteTendencyStruct
	}{
		// TODO: Add test cases.
		{"case 1", args{data, "petal-length", 2}, DiscreteTendencyStruct{
			ColName:    "petal-length",
			Max:        6.90,
			Min:        1.00,
			Range:      5.90,
			Variance:   3.09,
			Std:        1.76,
			Quantile25: 1.60,
			Quantile50: 4.30,
			Quantile75: 5.10,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiscreteTendency(tt.args.df, tt.args.colName, tt.args.decimalPlaces); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DiscreteTendency() = %v, want %v", got, tt.want)
			}
		})
	}
}
