package Plot

import (
	"github.com/go-gota/gota/dataframe"
	"goAnalysisi/DataSource"
	"testing"
)

func TestHistogramDataFrame(t *testing.T) {
	data := DataSource.ReadCSVDataFrame("../Resource/DataSet/iris-data.csv")
	type args struct {
		df         dataframe.DataFrame
		chooseCols []string
		width      int
		title      string
		SaveDict   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case 1", args{
			data, []string{"sepal-length", "sepal-width", "petal-length", "petal-width"}, 16, "iris-data", "../Resource/images/HistogramDataFrame/",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HistogramDataFrame(tt.args.df, tt.args.chooseCols, tt.args.width, tt.args.title, tt.args.SaveDict)
		})
	}
}

func TestBoxDataFrame(t *testing.T) {
	data := DataSource.ReadCSVDataFrame("../Resource/DataSet/iris-data.csv")
	type args struct {
		df         dataframe.DataFrame
		chooseCols []string
		size       float64
		title      string
		yLabel     string
		SaveDict   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case 1", args{
			data, []string{"sepal-length", "sepal-width", "petal-length", "petal-width"}, 50, "iris-data", "value", "../Resource/images/BoxDataFrame/",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BoxDataFrame(tt.args.df, tt.args.chooseCols, tt.args.size, tt.args.title, tt.args.yLabel, tt.args.SaveDict)
		})
	}
}
