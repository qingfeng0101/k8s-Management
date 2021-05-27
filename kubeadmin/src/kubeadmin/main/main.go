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







	//r.GET("/prod/namespace",prod.Namespace)
	//r.GET("/prod/nodes",prod.GetNode)
	//r.POST("/prod/pod",prod.Pod)
	//r.POST("/prod/deletepod",prod.DeletePod)
	//r.POST("/prod/getpodinfo",prod.GetPodInfo)
	//r.GET("/prod/getpodlog",prod.GetPodlog)
	//r.POST("/prod/getdeplyment",prod.Getdeployment)
	//r.POST("/prod/deletedeployment",prod.DeleteDeployment)
	//r.POST("/prod/updatadeployment",prod.UpdataDeployment)
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}