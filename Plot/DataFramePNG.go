package Plot

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/google/uuid"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// HistogramDataFrame 对DataFrame数据每个字段绘制直方图，保留png文件到指定路径/**
func HistogramDataFrame(df dataframe.DataFrame, chooseCols []string, width int, title string, SaveDict string) {
	dfsChose := df.Select(chooseCols)
	//log.Printf("\n%v", dfsChose)
	for _, colName := range dfsChose.Names() {
		values := make(plotter.Values, dfsChose.Nrow())
		for i, f := range dfsChose.Col(colName).Float() {
			values[i] = f
		}
		p := plot.New()
		if title != "" {
			p.Title.Text = fmt.Sprintf("%s", title)
		} else {
			p.Title.Text = fmt.Sprintf("%s", uuid.New().String())
		}
		hist, err := plotter.NewHist(values, width)
		if err != nil {
			log.Fatal(err)
		}
		hist.Normalize(1)
		p.Add(hist)
		if strings.HasSuffix(SaveDict, "/") {
			SaveDict = SaveDict[:len(SaveDict)-1]
		}
		dir, _ := os.Getwd()
		resPath := filepath.Join(dir, SaveDict, "/"+colName+"_hist.png")
		err = p.Save(4*vg.Inch, 4*vg.Inch, resPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// BoxDataFrame 对DataFrame数据所有字段绘箱形图，保留png文件到指定路径/**
func BoxDataFrame(df dataframe.DataFrame, chooseCols []string, size float64, title string, yLabel string, SaveDict string) {
	dfsChose := df.Select(chooseCols)
	p := plot.New()
	if title != "" {
		p.Title.Text = fmt.Sprintf("%s", title)
	} else {
		p.Title.Text = fmt.Sprintf("%s", uuid.New().String())
	}
	p.Y.Label.Text = yLabel
	w := vg.Points(size)
	for index, colName := range dfsChose.Names() {
		values := make(plotter.Values, dfsChose.Nrow())
		for i, f := range dfsChose.Col(colName).Float() {
			values[i] = f
		}
		boxPlot, err := plotter.NewBoxPlot(w, float64(index), values)
		if err != nil {
			log.Fatal(err)
		}
		p.Add(boxPlot)
	}

	p.NominalX(dfsChose.Names()...)
	if strings.HasSuffix(SaveDict, "/") {
		SaveDict = SaveDict[:len(SaveDict)-1]
	}
	dir, _ := os.Getwd()
	resPath := filepath.Join(dir, SaveDict, fmt.Sprintf("%s_BoxPlot.png", title))
	err := p.Save(4*vg.Inch, 4*vg.Inch, resPath)
	if err != nil {
		log.Fatal(err)
	}
}
