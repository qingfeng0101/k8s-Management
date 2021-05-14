package main

import "github.com/gin-gonic/gin"
import "kubeadmin/controller"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
		})
	r.GET("/namespace",controller.Namespace)
	r.POST("/pod",controller.Pod)
	r.POST("/deletepod",controller.DeletePod)
	r.POST("/getpodinfo",controller.GetPodInfo)
	r.GET("/getpodlog",controller.GetPodlog)
	r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}