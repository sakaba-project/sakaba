package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbConfig struct {
	user     string
	pass     string
	name     string
	protocol string
	params   string
}

func Connect() *gorm.DB {
	config := dbConfig{
		user: os.Getenv("MYSQL_USER"),
		pass: os.Getenv("MYSQL_PASSWORD"),
		name: os.Getenv("MYSQL_DATABASE"),
		protocol: fmt.Sprintf(
			"tcp(%s:%s)",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
		),
		params: fmt.Sprintf(
			"charset=%s&parseTime=%s&loc=%s",
			os.Getenv("DB_CHARSET"),
			os.Getenv("DB_PARSE_TIME"),
			os.Getenv("DB_LOC"),
		),
	}

	source := config.user + ":" + config.pass + "@" + config.protocol + "/" + config.name + "?" + config.params

	conn, err := gorm.Open("mysql", source)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
