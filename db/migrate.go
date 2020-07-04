package db

import (
	"github.com/sakaba-project/sakaba/models"
)

func AutoMigrate() {
	conn := Connect()
	defer conn.Close()

	conn.AutoMigrate(&models.User{})
	conn.Set("gorm:table_options", "ENGINE=InnoDB")
}
