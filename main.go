package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"zhuce/common"
)

func main() {
	db := common.InitDB()

	r := gin.Default()
	r = CollectRoute(r)

	log.Println(db)
	panic(r.Run()) // 监听并在 0.0.0.0:8080 上启动服务
}
