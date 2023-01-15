package service

import (
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"fmt"
)

func GetUserIdByToken(token string) uint {
	//通过传进来的token 获得userid
	db := ConnSql.ThemodelOfSql()
	var username string
	for _, v := range token {
		if v == '@' {
			break
		}
		username = fmt.Sprintf("%s%c", username, v) //进行字符串的拼接
	}
	println("通过token获取到的用户名是: ", username)
	var tmpUser Model.User
	db.Table("tik_user").Where("name = ?", username).First(&tmpUser)
	println("通过得到的用户名查询到的ID是", tmpUser.Id)
	return tmpUser.Id

}

//添加视频信息
func CreateVideo(video *Model.Video) {
	db := ConnSql.ThemodelOfSql()
	db.Table("videos").Create(&video)
}
