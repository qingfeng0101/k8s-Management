package controller

import (
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global/views"
)

func SSH(c *gin.Context)  {
	views.ShellWs(c)

}
