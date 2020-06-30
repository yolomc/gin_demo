package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloMsg(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "Hello World!"})
}

func main() {

	r := gin.Default()

	r.GET("/hello", helloMsg)

	r.Run(":9090")

}
