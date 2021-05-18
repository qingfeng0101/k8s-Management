package main

import "github.com/gin-gonic/gin"
import "kubeadmin/controller"

func main() {
	r := gin.Default()
	r.GET("/namespace",controller.Namespace)
	r.GET("/nodes",controller.GetNode)
	r.POST("/pod",controller.Pod)
	r.POST("/deletepod",controller.DeletePod)
	r.POST("/getpodinfo",controller.GetPodInfo)
	r.GET("/getpodlog",controller.GetPodlog)
	r.POST("/getdeplyment",controller.Getdeployment)
	r.POST("/deletedeployment",controller.DeleteDeployment)
	r.POST("/updatadeployment",controller.UpdataDeployment)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}