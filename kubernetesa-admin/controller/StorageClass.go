package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetStorageClassList(c *gin.Context)  {
	env := tools.GetEnv(c)
	StorageClasslist ,err := k8s.GetStorageClass(global.K8sclientset(env),"")
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"StorageClasslist":StorageClasslist,
	})

}
