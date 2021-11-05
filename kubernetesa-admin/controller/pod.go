package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetPodList(c *gin.Context)  {
	env := tools.GetEnv(c)
	namespace := c.Query("name")
	podlist ,err := k8s.GetPod(global.K8sclientset(env),"",namespace)
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"podlist":podlist,
	})

}
