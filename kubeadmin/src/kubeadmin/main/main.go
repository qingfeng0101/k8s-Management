package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubeadmin/controller"
)
func main() {
	r := gin.Default()
	ip := controller.Conf.Ip
	port := controller.Conf.Port
	ipprot := fmt.Sprintf("%s:%s", ip, port)
	r.POST("/api/namespace",controller.Namespace)
	r.POST("/api/nodes",controller.GetNode)
	r.POST("/api/pod",controller.Pod)
	r.POST("/api/deletepod",controller.DeletePod)
	r.POST("/api/getpodinfo",controller.GetPodInfo)
	r.POST("/api/getdeploymentinfo",controller.GetDeploymentInfo)
	r.GET("/api/getpodlog",controller.GetPodlog)
	r.POST("/api/getdeplyment",controller.Getdeployment)
	r.POST("/api/deletedeployment",controller.DeleteDeployment)
	r.POST("/api/updatadeployment",controller.UpdataDeployment)
	r.POST("/api/updateimages",controller.UpdataImages)
	r.POST("/api/upload",controller.Upload)
	r.GET("/api/names",controller.Getnames)
	r.POST("/api/delname",controller.Delnames)
	r.GET("/api/exec",controller.EXEC)
	r.Run(ipprot) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}