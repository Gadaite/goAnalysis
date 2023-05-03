package DataSource

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var MySqlDB *gorm.DB
var PgSqlDB *gorm.DB

const (
	mysqlUser        = "root"
	mysqlPassword    = "123456"
	mysqlHost        = "1.15.94.16"
	mysqlPort        = 3306
	mysqlDbname      = "mysqlcrud"
	postgresUser     = "postgres"
	postgresPassword = "zzjz123"
	postgresHost     = "1.15.94.16"
	postgresPort     = 5432
	postgresDbName   = "trajectory"
)

func InitMySql() {
	mysqlDsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDbname)
	db, err := gorm.Open("mysql", mysqlDsn)
	db.LogMode(true)
	db.SingularTable(true)
	if err != nil {
		fmt.Println("连接Mysql失败：" + err.Error())
	}
	MySqlDB = db
}

func InitPgSql() {
	db, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDbName))
	db.LogMode(true)
	db.SingularTable(true)
	if err != nil {
		panic("connect to postgresql failed :" + err.Error())
	}
	PgSqlDB = db
}
