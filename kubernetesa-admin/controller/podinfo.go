package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetPodInfo(c *gin.Context)  {
	env := tools.GetEnv(c)
	namespace := c.Query("namespace")
	podname := c.Query("podname")
	fmt.Println("podname: ",podname)
	fmt.Println("namespace: ",namespace)
	podinfo ,err := k8s.GetPodInfo(global.K8sclientset(env),namespace,podname)
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
		 return
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"podinfo":podinfo,
	})

}
