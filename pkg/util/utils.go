package util

import (
	"grpc_demo/conf"
	"strings"
)

// 获取mysql数据库dsn
func GetMysqlDSN() string {
	dsn := strings.Join([]string{conf.Mysql.Username, ":", conf.Mysql.Password, "@tcp(",
		conf.Mysql.Addr, ")/", conf.Mysql.Database, "?charset=" + conf.Mysql.Charset + "&parseTime=true"}, "")

	return dsn
}
