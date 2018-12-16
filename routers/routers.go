package routers

import (
  "github.com/gin-gonic/gin"
  "vq0599/controller"
)

func InitRouter() *gin.Engine {
  router := gin.Default()
  apiv1 := router.Group("/api/v1")

  apiv1.GET("/articles", controller.GetArticles)
  apiv1.GET("/articles/:id", controller.GetArticle)
  apiv1.POST("/articles", controller.AddArticle)
  apiv1.DELETE("/articles/:id",  controller.DeleteArticle)
  apiv1.PATCH("/articles/:id", controller.UpdateArticle)

  apiv1.POST("/login", controller.Login)
  return router
}