package routers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
    router := gin.Default()

    router.GET("api/v1/user/:id", func(c *gin.Context) {
        id := c.Params.ByName("id")
        c.JSON(http.StatusOK, gin.H{
            "id": id,
            "name": "黄努努",
            "age": 25,
        })
    })

    return router
}