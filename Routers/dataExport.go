package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	img "github.com/tmp_speed_server_go/Controller"
)

func main() {
	r := gin.Default()
	// 导出excel
	r.GET("/export", img.Export)
	r.Run() // listen and serve on 0.0.0.0:8080
}
