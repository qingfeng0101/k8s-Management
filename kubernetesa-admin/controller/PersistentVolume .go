package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetPersistentVolumeList(c *gin.Context)  {
	env := tools.GetEnv(c)
     pvlist ,err := k8s.GetPersistentVolume(global.K8sclientset(env),"",)
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"pvlist":pvlist,
	})

}
