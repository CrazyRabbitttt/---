package ConnSql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ThemodelOfSql() *gorm.DB {
	return db
}

func InitDb() {
	//数据库的连接
	dsn := "root:shaoguixin1+@tcp(127.0.0.1:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"

	db1, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		println("Error: ", err.Error())
	}
	db = db1 //将数据传送到全局变量中
	//进行表的创建
	//db.AutoMigrate(&Model.Video{})
	//db.AutoMigrate(&Model.Following{})
	//db.AutoMigrate(&Model.Favoriete{})
	//db.AutoMigrate(&Model.User{})

}
