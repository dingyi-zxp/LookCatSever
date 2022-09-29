package configs

type SqlConfigs struct {
	Host     int
	User     string
	Password string
	Address  string
}

func MysqlConfigs() SqlConfigs {
	var sqlConfigs SqlConfigs

	sqlConfigs.User = "root"
	sqlConfigs.Password = "cat020605"
	sqlConfigs.Host = 33777
	sqlConfigs.Address = "127.0.0.1"
	return sqlConfigs
}
