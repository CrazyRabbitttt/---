package controller

import (
	"Web-Go/Common"
	"Web-Go/ConnSql"
	"Web-Go/Model"
	"Web-Go/service"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type UserIdTokenResponse struct {
	UserId int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
	Common.Response
	UserIdTokenResponse
}

type UserLoginResponse struct {
	Common.Response
	UserIdTokenResponse
}

//用户注册的主函数， 最上层的接口函数
func UserRegister(c *gin.Context) {
	//传进来的参数的获取
	username := c.Query("username")
	password := c.Query("password")

	//进行service层次的处理
	registerResponse, err := UserRegisterService(username, password)

	//将响应进行返回
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			Response: Common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, UserRegisterResponse{
		Response:            Common.Response{StatusCode: 0},
		UserIdTokenResponse: registerResponse,
	})
	return
}

//用户进行登陆的处理函数，鉴别是否是存在的等
func UserRegisterService(userName string, passWord string) (UserIdTokenResponse, error) {

	var userResponse = UserIdTokenResponse{}

	//1.Legal check
	err := service.IsUserLegal(userName, passWord)
	if err != nil {
		return userResponse, err
	}
	//2.Create New User, 返回的对象中只是有用户名、密码
	var newUser Model.User
	newUser, err = service.CreateNewUser(userName, passWord)
	if err != nil {
		if err == Common.ErrorUserExits {
			//print("Error : user exist....")
			return userResponse, Common.ErrorUserExits //将err继续传递
		}
	}
	//进行token的颁发

	token := newUser.Name + newUser.Password + "bing"

	userResponse = UserIdTokenResponse{
		UserId: newUser.Id,
		Token:  token,
	}
	return userResponse, nil
}

//用户进行登陆的接口函数
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + password + "bing"
	userLoginResponse, err := UserLoginService(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Common.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	//如果用户是存在的话，返回对应的token 和 id
	userLoginResponse.Token = token
	c.JSON(http.StatusOK, UserLoginResponse{
		Response:            Common.Response{StatusCode: 0, StatusMsg: username + "登陆成功！"},
		UserIdTokenResponse: userLoginResponse,
	})
}

//用于提供检查等操作的辅助Login函数
func UserLoginService(userName string, passWrod string) (UserIdTokenResponse, error) {
	db := ConnSql.ThemodelOfSql()
	var userResponse = UserIdTokenResponse{}
	//进行数据的合法性检查
	err := service.IsUserLegal(userName, passWrod)
	if err != nil {
		return userResponse, err
	}

	//查询用户是否是存在的
	var tmpLoginUser Model.User

	result := db.Table("tik_user").Where("name = ?", userName).First(&tmpLoginUser)
	if result.Error != nil {
		//如果不存在记录的话就返回错误
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return userResponse, result.Error
		}
	}
	userResponse.UserId = tmpLoginUser.Id //将读取到的ID写会到response中
	return userResponse, nil
}

//
//func UserInfo(c *gin.Context) {
//	db := ConnSql.ThemodelOfSql()
//	userId := c.Query("user_id")
//
//	var dbUser User
//
//	fmt.Println("传入的user_id", userId)
//	db.Table("tik_user").Where("id = ?", userId).Find(&dbUser)
//
//	fmt.Println("查到的用户信息：", dbUser)
//
//	if dbUser.Id != 0 {
//		c.JSON(http.StatusOK, UserResponse{
//			Response: Response{Statuscode: 0, StatusMsg: "查询用户信息成功"},
//			User: User{
//				Id:            dbUser.Id,
//				Name:          dbUser.Name,
//				FollowCount:   188,
//				FollowerCount: 199,
//				IsFollow:      true,
//			},
//		})
//	} else {
//		c.JSON(http.StatusBadRequest, UserLoginResponse{
//			Response: Response{Statuscode: 1, StatusMsg: "查询失败"},
//			UserId:   dbUser.Id,
//		})
//	}
//
//}
