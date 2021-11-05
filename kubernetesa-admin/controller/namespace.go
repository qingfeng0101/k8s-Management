package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetNamespaceList(c *gin.Context)  {
	env := tools.GetEnv(c)
     namespaces ,err := k8s.GetNamespace(global.K8sclientset(env),"")
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"namespaces":namespaces,
	})

}
