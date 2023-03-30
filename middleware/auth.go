package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthToken(c *gin.Context){
	token := c.Request.Header.Get("Authorization")
	if token == ""{
		c.JSON(http.StatusBadRequest, gin.H{"error":"token not found"})
		c.Abort()
		return
	}
	fields := strings.Fields(token)
	if(fields[1] != "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiam95YWwgIiwiaWF0IjoxNjY1Nzc3MjEwLCJleHAiOjE2NjU3NzcyMjV9.NZ6uvufiCtjji4kYrhnqHL0vfFWZdwC6-htQU_it5f8"){
		c.JSON(http.StatusBadRequest,gin.H{"error":"token not valid"})
		c.Abort()
		return
	}
	c.Next()

}