package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/tools"
	"kubernetesa-admin/k8s"
	"net/http"
)

func GetNodeList(c *gin.Context)  {
	 env := tools.GetEnv(c)

     nodes,nodemetrics,memnum ,err := k8s.GetNode(global.K8sclientset(env),global.K8smetricsset(env),"")
     fmt.Println("k8s.GetNode err: ",err)
     if err != nil{
     	c.JSON(http.StatusOK,gin.H{
     		"code": 1,
     		"message": err,
		})
	 }
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"message": err,
		"nodes":nodes,
		"nodemetrics": nodemetrics,
		"memnum":memnum,
	})

}
