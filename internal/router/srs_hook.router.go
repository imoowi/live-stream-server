/*
generated by comer,https://github.com/imoowi/comer
Copyright © 2023 jun<simpleyuan@gmail.com>
*/
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/imoowi/live-stream-server/internal/middlewares"
	"github.com/imoowi/live-stream-server/internal/controllers"
)

func init() {
	RegisterRoute(SrsHookRouters)
}

func SrsHookRouters(e *gin.Engine) {
	api := e.Group("/api")
	api.Use(middlewares.JWTAuthMiddleware())
	api.Use(middlewares.CasbinMiddleware())
	api.Use(middlewares.UserLogMiddleware())
	srsHooks := api.Group("/srs-hooks")
	{
		srsHooks.GET("", controllers.SrsHookPageList) //分页
		srsHooks.GET("/:id", controllers.SrsHookOne) //一个
		srsHooks.POST("", controllers.SrsHookAdd) //新增
		srsHooks.PUT("/:id", controllers.SrsHookUpdate) //更新
		srsHooks.DELETE("/:id", controllers.SrsHookDel) //软删除
	}
}
