package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type bodyScheme struct {
  Name string `json:"name"`
  Email string `json:"email"`
}

func requestHandler (router *gin.Engine) {
  (*router).Any("/", func(ctx *gin.Context) {

    if ctx.Request.Method == "GET" {
      (*ctx).JSON(http.StatusOK, gin.H{
        "msg": "Success",
      })
    } 
    if ctx.Request.Method == "POST" {
      var user bodyScheme
      if err := ctx.ShouldBindJSON(&user); err != nil  {
        ctx.JSON(http.StatusBadRequest, gin.H{
          "msg" : "Failed",
        })
        return
      }

      ctx.JSON(http.StatusOK, gin.H{
        "msg": "Sucess",
        "data": user,
      })
    }
  })
}

func main() { 
  r := gin.Default()
  requestHandler(r)
  r.Run()
}