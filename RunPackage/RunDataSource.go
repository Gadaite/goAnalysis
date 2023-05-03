package RunPackage

import (
	"fmt"
	"goAnalysisi/DataSource"
	"time"
)

func CSVPart() {
	DataSource.ShowCSVArray("./Resource/DataSet/iris-data.csv")
	DataSource.ShowCSVDataFrame("./Resource/DataSet/iris-data.csv")
}

func SqlPart() {
	DataSource.ShowMysqlRows("select * from user")
	DataSource.ShowPgSqlRows("select * from user")
}

func MemoryCachePart() {
	DataSource.InitCacheMemory(5*time.Minute, 1*time.Minute)
	DataSource.Cache2Memory("yangyi", "123456789")
	for {
		cache := DataSource.LoadMemoryCache("yangyi")
		fmt.Println(cache)
		time.Sleep(20 * time.Second)
	}
}

func BoltCachePart() {
	DataSource.InitCacheBolt("./Resource/Bolt/test.db", 0600, nil)
	DataSource.CreateBoltBucket("QuickStart")
	DataSource.Cache2Bucket("yangyi", "987654321", "QuickStart")
	DataSource.LoadBucketCache("QuickStart")
}
