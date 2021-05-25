package main

import (
	"github.com/gin-gonic/gin"
    "kubeadmin/controller"
	"kubeadmin/prod"
)
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
	r.GET("/prod/namespace",prod.Namespace)
	r.GET("/prod/nodes",prod.GetNode)
	r.POST("/prod/pod",prod.Pod)
	r.POST("/prod/deletepod",prod.DeletePod)
	r.POST("/prod/getpodinfo",prod.GetPodInfo)
	r.GET("/prod/getpodlog",prod.GetPodlog)
	r.POST("/prod/getdeplyment",prod.Getdeployment)
	r.POST("/prod/deletedeployment",prod.DeleteDeployment)
	r.POST("/prod/updatadeployment",prod.UpdataDeployment)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}