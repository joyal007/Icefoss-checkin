package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joyal007/go-scan/controllers"
)

func Routes(routes *gin.RouterGroup){

	routes.POST("/culturals",controllers.Culturals)
	routes.GET("/culturals",controllers.GetCulturals)

	routes.POST("/workshop/go",controllers.GoWorkshop)
	routes.POST("/workshop/devops",controllers.DevopsWorkshop)
	routes.POST("/hackfit",controllers.Hackfit)
	routes.POST("/workshop/flutter",controllers.FlutterWorkshop)
	routes.POST("/workshop/webdev",controllers.WebdevWorkshop)
	routes.POST("/workshop/rust",controllers.RustWorkshop)
	routes.POST("/workshop/ar",controllers.ArWorkshop)

	routes.GET("/hackfit",controllers.HackfitCheckIn)
	routes.GET("/workshop/go",controllers.GoWorkshopCheckIn)
	routes.GET("/workshop/devops",controllers.DevopsWorkshopCheckIn)
	routes.GET("/workshop/flutter",controllers.FlutterWorkshopCheckIn)
	routes.GET("/workshop/webdev",controllers.WebWorkshopCheckIn)
	routes.GET("/workshop/rust",controllers.RustWorkshopCheckIn)
	routes.GET("/workshop/ar",controllers.ArWorkshopCheckIn)
}