package controller

import (
	"Web-Go/Common"
	"Web-Go/Model"
	"Web-Go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"path/filepath"
	"strconv"
)

type VideoListResponse struct {
	Common.Response
	VideoList []ReturnVideo `json:"video_list"`
}

type ReturnAuthor struct {
	AuthorId      uint   `json:"author_id"`
	Name          string `json:"name"`
	FollowCount   uint   `json:"follow_count"`
	FollowerCount uint   `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}

type ReturnVideo struct {
	VideoId       uint         `json:"video_id"`
	Author        ReturnAuthor `json:"author"`
	PlayUrl       string       `json:"play_url"`
	CoverUrl      string       `json:"cover_url"`
	FavoriteCount uint         `json:"favorite_count"`
	CommentCount  uint         `json:"comment_count"`
	IsFavorite    bool         `json:"is_favorite"`
	Title         string       `json:"title"`
}

//将视频上传到本地的服务器上面
func Publish(c *gin.Context) {
	//1. 通过token 获得用户名
	token := c.PostForm("token")
	userId := service.GetUserIdByToken(token)
	println("通过token获取到的用户id：", userId)
	title := c.PostForm("title")
	println("获取到的title:", title)
	//2.进行传入视频文件位置的设置
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	filename := filepath.Base(data.Filename)            //去除掉路径的文件名称
	finalName := fmt.Sprintf("%d_%s", userId, filename) //最终存储的文件名称
	saveFile := filepath.Join("./public/", finalName)   //拼接成为绝对路径
	println("The save path is :", saveFile)

	if err := c.SaveUploadedFile(data, saveFile); err != nil { //进行文件的存储
		c.JSON(http.StatusOK, Common.Response{
			StatusCode: 1,
			StatusMsg:  err.Error() + "文件存储失败",
		})
		return
	}
	c.JSON(http.StatusOK, Common.Response{
		StatusCode: 0,
		StatusMsg:  finalName + " upload successfully",
	})

	//进行video的信息的配置
	playUrl := "http://192.168.43.104:8080/" + "public/" + finalName
	coverUrl := "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"

	video := Model.Video{
		Model:         gorm.Model{},
		AuthorId:      uint(userId),
		PlayUrl:       playUrl,
		CoverUrl:      coverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		Title:         title,
	}
	service.CreateVideo(&video)
}

//获取视频列表的方式
func PublishList(c *gin.Context) {
	//传入的是token(当前用户) & userid(客人用户)
	token := c.Query("token")

	hostId := service.GetUserIdByToken(token) //hostId
	s_userId := c.Query("user_id")
	guestId, err := strconv.ParseUint(s_userId, 10, 64)
	if err != nil {
		//do nothing
		println("将userid -> uint 失败")
	}
	//通过userid进行用户信息的查询
	getUser, err := service.GetUserById(uint(guestId))
	if err != nil {
		c.JSON(http.StatusOK, Common.Response{
			StatusCode: 1,
			StatusMsg:  "Not found the person",
		})
		return
	}

	returnAuthor := ReturnAuthor{
		AuthorId:      uint(guestId),
		Name:          getUser.Name,
		FollowCount:   getUser.FollowCount,
		FollowerCount: getUser.FollowerCount,
		IsFollow:      service.IsFollowing(hostId, uint(guestId)),
	}

	videoList := service.GetVideoListById(uint(guestId))
	if len(videoList) == 0 {
		c.JSON(http.StatusOK, VideoListResponse{
			Response:  Common.Response{StatusCode: 1, StatusMsg: "用户的视频数目是0"},
			VideoList: nil,
		})
	} else { //需要展示的资源进行返回
		var returnVideoList []ReturnVideo
		for i := 0; i < len(videoList); i++ {
			returnVideo := ReturnVideo{
				VideoId:       videoList[i].ID,
				Author:        returnAuthor,
				PlayUrl:       videoList[i].PlayUrl,
				CoverUrl:      videoList[i].CoverUrl,
				FavoriteCount: videoList[i].FavoriteCount,
				CommentCount:  videoList[i].CommentCount,
				Title:         videoList[i].Title,
				IsFavorite:    service.CheckFavorite(hostId, videoList[i].ID),
			}
			//将video追加到最终的videoList中去
			returnVideoList = append(returnVideoList, returnVideo)
		}
		c.JSON(http.StatusOK, VideoListResponse{
			Response: Common.Response{
				StatusCode: 0,
				StatusMsg:  "successfully got the VideoList",
			},
			VideoList: returnVideoList,
		})
	}

}
