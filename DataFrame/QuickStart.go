package DataFrame

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	_ "github.com/go-gota/gota/series"
	_ "log"
	"os"
)

func QuickStart()  {
	// 读取 CSV 文件
	file, _ := os.Open("F:\\Code\\DataSet\\MLdataset\\adult.csv")
	df := dataframe.ReadCSV(file)
	fmt.Println(df)
}
