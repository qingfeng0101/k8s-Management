package main

import (
	"github.com/gin-gonic/gin"
	"kubeadmin/controller"
)
func main() {
	r := gin.Default()
	r.POST("/api/namespace",controller.Namespace)
	r.POST("/api/nodes",controller.GetNode)
	r.POST("/api/pod",controller.Pod)
	r.POST("/api/deletepod",controller.DeletePod)
	r.POST("/api/getpodinfo",controller.GetPodInfo)
	r.GET("/api/getpodlog",controller.GetPodlog)
	r.POST("/api/getdeplyment",controller.Getdeployment)
	r.POST("/api/deletedeployment",controller.DeleteDeployment)
	r.POST("/api/updatadeployment",controller.UpdataDeployment)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}