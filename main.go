package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joyal007/go-scan/routes"
)


func main(){

	PORT:=os.Getenv("PORT")
	if PORT ==""{
		PORT = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	api := router.Group("icefoss")
	v1 := api.Group("v1")
	checkin := v1.Group("checkin")

	routes.Routes(checkin)

	router.Run(":"+PORT)
}