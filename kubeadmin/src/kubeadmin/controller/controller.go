package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"kubeadmin/model"
	"kubeadmin/pulick"
	"net/http"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Namespace(c *gin.Context)  {
	ENV := c.PostForm("env")
	switch  {
	case ENV == "test":
		namespacelist,err:=pulick.NamespaceList(clientset)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,namespacelist)
	case ENV == "prod":
		namespacelist,err:=pulick.NamespaceList(prodclientset)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,namespacelist)
	}
}
func Pod(c *gin.Context)  {
	ENV := c.PostForm("env")
	ns := c.PostForm("namespace")
	switch  {
	case ENV == "test":
		pods,err:=pulick.Pods(clientset,ns)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,pods)
	case ENV == "prod":
		pods,err:=pulick.Pods(prodclientset,ns)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,pods)
	}


}

func DeletePod(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	switch  {
	case ENV == "test":
		OK,err:=pulick.DeletePod(clientset,name,namespace)
		if err != nil{
			fmt.Println("DeletePod err: ",err)
			return
		}
		if OK{
			// 删除成功后查询返回给前端
			pods,err := pulick.Pods(clientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,pods)
		}
	case ENV == "prod":
		OK,err:=pulick.DeletePod(prodclientset,name,namespace)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		if OK{
			pods,err := pulick.Pods(prodclientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,pods)
		}
	}


}
func GetPodInfo(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	switch  {
	case ENV == "test":
		podinfo,err:=pulick.PodINFO(clientset,name,namespace)
		if err != nil{
			fmt.Println("PodINFO err: ",err)
			return
		}
		pulick.Get(c,podinfo)
	case ENV == "prod":
		podinfo,err:=pulick.PodINFO(prodclientset,name,namespace)
		if err != nil{
			fmt.Println("PodINFO err: ",err)
			return
		}
		pulick.Get(c,podinfo)
	}


}
func GetPodlog(c *gin.Context)  {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		 var mess  model.Message
		//var d map[string]interface{}
		// var i interface{}
		err = json.Unmarshal(message, &mess)
		if err != nil{
			fmt.Println("Unmarshal err: ",err)
			return
		}
		fmt.Println(mess)
		switch {
		case mess.ENV == "test":
			pulick.PodLogs(ws, mt, clientset, mess.Name, mess.Namespace)
		case mess.ENV == "prod":
			pulick.PodLogs(ws, mt, prodclientset, mess.Name, mess.Namespace)
		}
	}
}
func Getdeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	ns := c.PostForm("namespace")
	switch  {
	case ENV == "test":
		DPList,err:=pulick.Deploymnet(clientset,ns)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,DPList)
	case ENV == "prod":
		DPList,err:=pulick.Deploymnet(prodclientset,ns)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,DPList)
	}

}
func DeleteDeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	switch  {
	case ENV == "test":
		OK,err:=pulick.DeleteDeployment(clientset,name,namespace)
		if err != nil{
			fmt.Println("DeletePod err: ",err)
			return
		}
		if OK{
			// 删除成功后查询返回给前端
			deployments,err := pulick.Deploymnet(clientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,deployments)
		}
	case ENV == "prod":
		OK,err:=pulick.DeleteDeployment(prodclientset,name,namespace)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		if OK{
			deployments,err := pulick.Deploymnet(prodclientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,deployments)
		}
	}

}
func UpdataDeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	num := c.PostForm("num")
	n,_ := strconv.Atoi(num)
	dnum := int32(n)
	// 扩缩容操作
	switch  {
	case ENV == "test":
		OK,err:=pulick.UpDeployment(clientset,name,namespace,&dnum)
		if err != nil{
			fmt.Println("DeletePod err: ",err)
			return
		}
		if OK{
			// 扩缩容成功后查询返回给前端
			deployments,err := pulick.Deploymnet(clientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,deployments)
		}
	case ENV == "prod":
		OK,err:=pulick.UpDeployment(prodclientset,name,namespace,&dnum)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		if OK{
			deployments,err := pulick.Deploymnet(prodclientset,namespace)
			if err != nil{
				fmt.Println("Pods err: ",err)
				return
			}
			pulick.Get(c,deployments)
		}
	}
}
func GetNode(c *gin.Context)  {
	ENV := c.PostForm("env")
	switch  {
	case ENV == "test":
		nodes,err:=pulick.Nodes(clientset)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,nodes)
	case ENV == "prod":
		nodes,err:=pulick.Nodes(prodclientset)
		if err != nil{
			fmt.Println("namespace err: ",err)
			return
		}
		pulick.Get(c,nodes)
	}
}
