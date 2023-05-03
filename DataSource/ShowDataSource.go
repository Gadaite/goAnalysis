package DataSource

import (
	"database/sql"
	"fmt"
	"log"
)

func ShowCSVArray(path string) {
	readCSV := ReadCSV2Array(path)
	for i := range readCSV {
		for j := range readCSV[i] {
			fmt.Print(readCSV[i][j] + " ")
		}
		fmt.Println()
	}
}

func ShowCSVDataFrame(path string) {
	frame := ReadCSVDataFrame(path)
	fmt.Println(frame)
}

func ShowMysqlRows(sql string) {
	rows := ReadMySqlRows(sql)
	printRows(rows)
	defer rows.Close()
}

func ShowPgSqlRows(sql string) {
	rows := ReadPgSqlRows(sql)
	printRows(rows)
	defer rows.Close()
}

func printRows(rows *sql.Rows) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	var values []interface{}
	for rows.Next() {
		record := make(map[string]interface{})
		values = make([]interface{}, len(columns))
		for i := range columns {
			values[i] = new(interface{})
			record[columns[i]] = nil
		}

		if err := rows.Scan(values...); err != nil {
			log.Fatal(err)
		}

		for i := range columns {
			switch val := (*(values[i].(*interface{}))).(type) {
			case []byte:
				record[columns[i]] = string(val)
			default:
				record[columns[i]] = val
			}
		}

		fmt.Printf("%v\n", record)
	}
}
