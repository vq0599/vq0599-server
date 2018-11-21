package routers

import (
    "github.com/gin-gonic/gin"
    "vq0599/controller"
)

func InitRouter() *gin.Engine {
    router := gin.Default()
    apiv1 := router.Group("/api/v1")

    apiv1.GET("/articles", controller.GetArticle)
    apiv1.POST("/articles/add", controller.AddArticle)

    return router
}