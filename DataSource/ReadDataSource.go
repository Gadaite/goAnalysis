package DataSource

import (
	"database/sql"
	"encoding/csv"
	"github.com/go-gota/gota/dataframe"
	"log"
	"os"
)

func ReadCSV2Array(path string) [][]string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}
	reader.FieldsPerRecord = -1
	var raw [][]string
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		raw = append(raw, record)
	}
	return raw
}

func ReadCSVDataFrame(path string) dataframe.DataFrame {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	df := dataframe.ReadCSV(file)
	return df
}

func ReadMySqlRows(sql string) *sql.Rows {
	InitMySql()
	rows, err := MySqlDB.DB().Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}

func ReadPgSqlRows(sql string) *sql.Rows {
	InitPgSql()
	rows, err := PgSqlDB.DB().Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
