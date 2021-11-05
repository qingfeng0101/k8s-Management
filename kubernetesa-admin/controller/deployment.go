package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetDeploymentList(c *gin.Context)  {
	env := tools.GetEnv(c)
	namespace := c.Query("name")
     deploymentlist ,err := k8s.Getdeployment(global.K8sclientset(env),"",namespace)
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"deploymentlist":deploymentlist,
	})

}
