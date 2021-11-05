package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"kubernetesa-admin/global"
	"net/http"
)

// 上传config文件接口
func Upload(c *gin.Context)  {
	// 获取集群名称
	Cluster_name := c.PostForm("ClusterName")
	// 获取前端传来的文件内容
	file := c.PostForm("files")
	fmt.Println("file: ",file)
	dst := fmt.Sprintf("%s/%s", global.Conf.Path, Cluster_name)
	// 字符串转换成字节数组
	filebyte := []byte(file)
	// 写入到本地
	err := ioutil.WriteFile(dst,filebyte,0666)
	if err != nil{
		fmt.Println("ioutil.WriteFile err: ",err)
		c.JSON(http.StatusInternalServerError,gin.H{
			"code": 1,
		})
		return
	}
	global.Init()
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
	})
	//c.SaveUploadedFile(file, dst)
	//Addclientsetmap(name,dst)
	//Getnames(c)

}
