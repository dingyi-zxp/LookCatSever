package linkSql

import (
	sqlInfo "LookCat/configs"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	var dsn = SpliceDsn()

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		fmt.Printf("Open dataset error: %v", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Printf("连接数据库失败,err:%v\n", err)
		return
	}
	fmt.Printf("success")

	str := db.QueryRow("SELECT * from other ")

	var (
		name string
		sex  string
	)

	str.Scan(&name, &sex)

	fmt.Println(name, sex)
}

func SpliceDsn() string {
	var info sqlInfo.SqlConfigs
	info = sqlInfo.MysqlConfigs()

	var dsn string

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/lookcat?charset=utf8", info.User, info.Password, info.Address, info.Host)
	fmt.Println(dsn)
	return dsn
}
