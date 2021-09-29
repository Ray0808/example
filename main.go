package main

import (
	controller "firstWeb/controllers"
	db "firstWeb/db"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var client *redis.Client

func main() {

	db.RedisConnect()

	r := gin.Default()
	// request := make(chan func)
	// receive := make(chan func)
	// go controller.SetRedis(){

	// }
	r.GET("/sendData", controller.SetRedis)
	r.GET("/getData", controller.GetRedis)
	r.GET("/editData", controller.ReWriteData)
	r.
		r.Run(":8080")
}
