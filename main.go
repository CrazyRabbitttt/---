package main

import (
	"Web-Go/ConnSql"
	"Web-Go/route"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default() //创建gin路由

	ConnSql.InitDb()

	route.InitRouter(r) //进行不同路由的处理

	r.Run() //Listen and serve on 0.0.0.0:8080

}
