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
	"strconv"
)

type UserIdTokenResponse struct {
	UserId uint   `json:"user_id"`
	Token  string `json:"token"`
}

type UserRegisterResponse struct {
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

	token := newUser.Name + "@" + newUser.Password + "bing"

	userResponse = UserIdTokenResponse{
		UserId: newUser.Id,
		Token:  token,
	}
	return userResponse, nil
}

type UserLoginResponse struct {
	Common.Response
	UserIdTokenResponse
}

//用户进行登陆的接口函数
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	token := username + "@" + password + "bing"
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
func UserLoginService(userName string, passWord string) (UserIdTokenResponse, error) {
	db := ConnSql.ThemodelOfSql()
	var userResponse = UserIdTokenResponse{}
	//进行数据的合法性检查
	err := service.IsUserLegal(userName, passWord)
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

//用户返回信息的结构体
type UserInfoQueryResponse struct {
	Id            uint   `json:"id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type UserInfoResponse struct {
	Common.Response
	UserList UserInfoQueryResponse `json:"user"`
}

//UserInfo 用户信息的主函数，传入：user_id token
func UserInfo(c *gin.Context) {
	userID := c.Query("user_id")

	userInfoResponse, err := UserInfoService(userID)

	if err != nil {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Common.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}

	//如果用户是存在的，返回对应的ID和token
	c.JSON(http.StatusOK, UserInfoResponse{
		Response: Common.Response{
			StatusCode: 0,
			StatusMsg:  "登陆成功",
		},
		UserList: userInfoResponse,
	})

}

// UserInfoService 进行用户信息的处理的函数
func UserInfoService(userID string) (UserInfoQueryResponse, error) {
	db := ConnSql.ThemodelOfSql()
	var tmpUserResponse UserInfoQueryResponse
	//将string转为int类型
	userId, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return tmpUserResponse, err
	}

	//下面进行数据的读取
	var tmpuser Model.User
	db.Table("tik_user").Where("id = ?", userId).First(&tmpuser)

	println(tmpuser.Id, tmpuser.Name, tmpuser.FollowCount)

	tmpUserResponse = UserInfoQueryResponse{
		Id:            tmpuser.Id,
		Name:          tmpuser.Name,
		FollowCount:   tmpuser.FollowCount,
		FollowerCount: tmpuser.FollowerCount,
		IsFollow:      false,
	}
	return tmpUserResponse, nil
}
