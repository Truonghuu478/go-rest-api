package routers

import "github.com/gin-gonic/gin"

func CartRoute(router *gin.Engine) {
	//All routes related to users comes here
	router.GET("/product/search")
}
