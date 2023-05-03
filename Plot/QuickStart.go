package Plot

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	gonumDraw "gonum.org/v1/plot/vg/draw"
	"gonum.org/v1/plot/vg/vgimg"
	_ "image"
	_ "image/color"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
)

func QuickStart1()  {
	getwd, _ := os.Getwd()
	p:= plot.New()
	line := plotter.XYs{
		{0, 0},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 7},
	}
	newLine, err := plotter.NewLine(line)
	if err ==nil {
		p.Add(newLine)
		length := vg.Length(500)
		width := vg.Length(500)
		canvas := vgimg.New(length, width)
		c := gonumDraw.New(canvas)
		p.Draw(c)
		err := p.Save(4*vg.Inch, 4*vg.Inch, filepath.Join(getwd,"Resource\\images\\lineplot.png"))
		if err!=nil {
			panic(err)
		}
	}
}

func QuickStart2(){
	// 创建一个新的 Plot
	p := plot.New()
	// 创建一个 Line
	line := plotter.XYs{
		{0, 0},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 7},
	}
	newLine, err := plotter.NewLine(line)
	if err != nil {
		log.Fatal(err)
	}

	// 将 Line 添加到 Plot 中
	p.Add(newLine)

	// 创建一个 500x500 的图像
	width := vg.Length(500)
	height := vg.Length(500)
	c := vgimg.New(width, height)

	// 创建 draw.Canvas 并将 plot.Plot 对象添加到其中
	dc := gonumDraw.New(c)
	p.Draw(dc)
}








