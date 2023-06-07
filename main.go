package main

import (
	"net/http"
  
	"github.com/gin-gonic/gin"
  )

func main(){
	r := gin.Default()//ginエンジン起動
	r.LoadHTMLGlob("templates/*")//HTMLをロード
	r.GET("/nakarisa", func(c *gin.Context) {//url
	  c.HTML(http.StatusOK, "index.html", gin.H{
		"message": "pong",
	  })
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}