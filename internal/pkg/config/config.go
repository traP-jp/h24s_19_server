package config

import (
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func getEnv(key, defaultValue string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}

	return v
}

func AppAddr() string {
	return getEnv("APP_ADDR", ":8080")
}

func MySQL() *mysql.Config {
	c := mysql.NewConfig()

	c.DBName = getEnv("NS_MARIADB_DATABASE", "database")
	c.User = getEnv("NS_MARIADB_USER", "user")
	c.Net = "tcp"
	c.Passwd = getEnv("NS_MARIADB_PASSWORD", "password")
	c.Addr = fmt.Sprintf(
		"%s:%s",
		getEnv("NS_MARIADB_HOSTNAME", "localhost"),
		getEnv("NS_MARIADB_PORT", "3306"),
	)
	c.Collation = "utf8mb4_general_ci"
	c.AllowNativePasswords = true
	return c
}
