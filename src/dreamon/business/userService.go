package business

import (
	"dreamon/config"
	"dreamon/models/mysql"
	"fmt"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	dbTemp, _ := config.LoadDB()
	if dbTemp != nil {
		db = dbTemp
		//先这样创建
		if !db.HasTable(&mysql.User{}) {
			db.CreateTable(&mysql.User{})
		}
	}
}
func AddUser(name string, emial string, password string) bool {
	user := mysql.User{
		Name:     name,
		Email:    emial,
		Password: password,
	}
	if !db.HasTable(&mysql.User{}) {
		fmt.Print("没有表")
	}
	if !db.NewRecord(&user) {
		return false
	} // => 主键为空返回`true`

	db.Create(&user)
	if !db.NewRecord(&user) {
		return true
	}
	return false
}
