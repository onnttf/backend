package dal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = ""

var DB *gorm.DB

func InitMySQL() error {
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
