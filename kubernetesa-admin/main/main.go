package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"kubernetesa-admin/global"
	"kubernetesa-admin/routes"
)





func main()  {
	// 初始化clientset
	if err := global.Init();err != nil{
		panic(err)
	}
	server := gin.Default()
	ip := global.Conf.Ip
	port := global.Conf.Port
	ipprot := fmt.Sprintf("%s:%s", ip, port)
    routes.Route(server)
	if err := server.Run(ipprot);err != nil{
		panic(err)
	}
}
