package tools

import "github.com/gin-gonic/gin"
var env string
func GetEnv(c *gin.Context) string {
	env := c.Query("env")
	return env
}