package configs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DB struct {
	Conn Conn `json:"conn"`
}

type Conn struct {
	Host     int    `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
}

func MysqlConfigs() Conn {
	var sqlConfigs Conn
	sqlConfigs = readFromJson()
	sqlConfigs.User = sqlConfigs.User
	sqlConfigs.Password = sqlConfigs.Password
	sqlConfigs.Host = sqlConfigs.Host
	sqlConfigs.Address = sqlConfigs.Address
	return sqlConfigs
}

func readFromJson() Conn {
	fileContent, err := os.Open("configs/Json/mysqlConfigs.json")
	checkError(err)
	fmt.Println("The File is opened successfully...")
	defer fileContent.Close()

	byteResult, _ := ioutil.ReadAll(fileContent)

	var sqlConfig DB
	json.Unmarshal(byteResult, &sqlConfig)

	fmt.Println(sqlConfig.Conn)

	return sqlConfig.Conn

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
