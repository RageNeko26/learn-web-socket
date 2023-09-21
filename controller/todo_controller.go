package controller

import (
	"learn-socket/config"
	"net/http"
  "fmt"

	"github.com/gin-gonic/gin"

)

var DataDummy = []string{"Apple", "Foo", "Bar"}

func GetHome(ctx *gin.Context) {
    ctx.JSON(200, gin.H{
      "message": "Todo API Websocket Version 1.0.0",
      "status": "success",
    })
}

func WSHandler(ctx *gin.Context) {
  config.Upgrader.CheckOrigin = func(r *http.Request) bool { return true }
  conn, err := config.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)

  if err != nil {
    return
  }

  defer conn.Close()

  for {
    messageType, p, err := conn.ReadMessage()

    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println(string(p))

    if err := conn.WriteMessage(messageType, []byte("Data from server...")); err != nil {
      fmt.Println(err)
      return
    }
  } 


}
