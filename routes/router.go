package routes

import (
	v1 "gin-casbin/api/v1"
	"gin-casbin/global"
	"gin-casbin/middleware"
	"gin-casbin/utils/load"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	global.GVA_Set = load.InitSet()
	gin.SetMode(global.GVA_Set.AppMode)
	router := gin.New()
	router.Use(middleware.Logger())
	userApiApp := &v1.UserApi{}
	router.Use(gin.Recovery())
	router.Use(middleware.Cors())
	publicGroup := router.Group("")
	{
		publicGroup.POST("register", userApiApp.Register)
	}
	privateGroup := router.Group("")
	privateGroup.Use(middleware.CasBin())
	privateGroup.Use(middleware.JwtToken())
	{
		privateGroup.GET("user/:uid", userApiApp.GetUserInfo)

	}
	err := router.Run(global.GVA_Set.HttpPort)
	if err != nil {
		return
	}
}
