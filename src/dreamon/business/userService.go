package business

import (
	"dreamon/config"
	"dreamon/models/mysql"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func init() {
	dbTemp, _ := config.OpenConnection()
	if dbTemp != nil {
		db = dbTemp
		//先这样创建
		// if !db.HasTable(&mysql.User{}) {
		// 	db.CreateTable(&mysql.User{})
		// }
	}
}
func AddUser(name string, mail string, password string) bool {
	user := mysql.User{
		Name:     name,
		Email:    mail,
		Password: password,
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

func GetUser(mail string, password string) *mysql.User {
	var user mysql.User
	db.Where(&mysql.User{Email: mail, Password: password}).First(&user)
	return &user
}
