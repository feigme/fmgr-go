package bootstrap

import (
	"github.com/feigme/fmgr-go/global"
	"github.com/feigme/fmgr-go/routes"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册 api 分组路由
	apiRoot := router.Group("/")
	routes.SetApiGroupRoutes(apiRoot)

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()
	r.Run(":" + global.App.Config.ApiServer.Port)
}
