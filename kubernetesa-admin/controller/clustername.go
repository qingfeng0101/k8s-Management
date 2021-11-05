package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func GetCluster(c *gin.Context)  {
	var ClusterNames []string

	filepath.Walk("/Users/zhouhao/down", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//fmt.Println("path: ",path)
			//env.Name = info.Name()
			//env.Url = info.Name()
			//names = append(names, env)
			fmt.Println("name: ", info.Name())
			ClusterNames = append(ClusterNames, info.Name())
		}


		return nil
	})
	c.JSON(http.StatusOK, gin.H{
		"code":         0,
		"ClusterNames": ClusterNames,
	})
    // if err != nil{
    // 	c.JSON(http.StatusOK,gin.H{
    // 		"code": 1,
    // 		"message": err,
	//	})
	// }
	//c.JSON(http.StatusOK,gin.H{
	//	"code": 0,
	//	"message": err,
	//	"nodes":nodes,
	//	"nodemetrics": nodemetrics,
	//	"memnum":memnum,
	//})

}
