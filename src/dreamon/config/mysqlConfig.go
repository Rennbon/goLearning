package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OpenConnection() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:qwe123456@/godb?charset=utf8&parseTime=True")
	//defer db.Close()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %s ", err)
	}
	db.DB().SetMaxIdleConns(10)
	return db, nil
}
