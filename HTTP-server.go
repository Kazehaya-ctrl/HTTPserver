package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  router := gin.Default()
  router.LoadHTMLGlob("./index.html")
  router.GET("/", func(ctx *gin.Context) {
    // data, err := os.ReadFile("index.html")
    // if err != nil {
    //   return
    // }
    data2 := gin.H{
      "msg" : "Sucess",
    }
    ctx.HTML(http.StatusOK, "index.html", data2)
  })
  // fmt.Println(string(data))
  router.Run()
}