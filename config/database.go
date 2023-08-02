package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQLConnection(cfg *MySQL) (*gorm.DB, error) {
	option := &gorm.Config{}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
	gormDB, err := gorm.Open(mysql.Open(dsn), option)
	if err != nil {
		return nil, err
	}
	message := fmt.Sprintf("successfully connect to %s", cfg.Name)
	fmt.Println(message)
	return gormDB, nil
}
