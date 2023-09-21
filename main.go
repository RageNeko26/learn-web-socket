package main

import (
	"learn-socket/controller"

	"github.com/gin-gonic/gin"	
)



func main() {
  r := gin.Default()

  r.GET("/", controller.GetHome)
  r.GET("/ws", controller.WSHandler)

  r.Run(":8080")
}
