package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect() *gorm.DB {
	const DBMS = "mysql"

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	protocol := fmt.Sprintf(
		"tcp(%s:%s)",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)
	params := fmt.Sprintf(
		"charset=%s&parseTime=%s&loc=%s",
		os.Getenv("DB_CHARSET"),
		os.Getenv("DB_PARSE_TIME"),
		os.Getenv("DB_LOC"),
	)

	connectedArgs := user + ":" + pass + "@" + protocol + "/" + dbName + "?" + params

	conn, err := gorm.Open(DBMS, connectedArgs)
	if err != nil {
		panic(err.Error())
	}

	return conn
}
