package cienv

import (
	"fmt"
	"os"
)

func GetEnv(name string) string {
	return os.Getenv(name)
}

func GetMySqlConfig(dbName string) string {
	host := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_HOST_WRITE", dbName))
	port := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_PORT", dbName))
	name := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_NAME", dbName))
	user := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_USERNAME", dbName))
	pass := GetEnv(fmt.Sprintf("MYSQL_DATABASE_%s_PASSWORD", dbName))
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, name)
}

func GetService(serviceKey string) string {
	return GetEnv(fmt.Sprintf("DEPLOYMENT_%s_HOST", serviceKey))
}
