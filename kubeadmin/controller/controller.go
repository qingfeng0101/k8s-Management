package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"kubeadmin/model"
	"kubeadmin/pulick"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Namespace(c *gin.Context)  {
	ENV := c.PostForm("env")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	namespacelist,err:=pulick.NamespaceList(clientset)
	if err != nil{
		fmt.Println("namespace err: ",err)
		return
	}
	pulick.Get(c,namespacelist)

}
func Pod(c *gin.Context)  {
	ENV := c.PostForm("env")
	ns := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	pods,err:=pulick.Pods(clientset,ns)
	if err != nil{
		fmt.Println("namespace err: ",err)
		return
	}
	pulick.Get(c,pods)
}

func DeletePod(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
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
}
func GetPodInfo(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	podinfo,err:=pulick.PodINFO(clientset,name,namespace)
	if err != nil{
		fmt.Println("PodINFO err: ",err)
		return
	}
	pulick.Get(c,podinfo)
}
func GetPodlog(c *gin.Context)  {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	statsus := make(chan bool)
	//ctx, stop := context.WithCancel(context.Background())
	for {

		mt, message, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("ReadMessage: ",err)
			fmt.Println("################################################")
			statsus <- false
			//stop()
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
		fmt.Println("cname: ",mess.Cname)
		fmt.Println("name: ",mess.Name)
		fmt.Println("第一层： ",mess.Status)
		fmt.Println(clientsetmap)
		clientset,OK := clientsetmap[mess.ENV]
		if !OK{
			fmt.Println("当前环境不存在")
			return
		}
		go pulick.PodLogs(ws, mt, clientset, mess.Name,mess.Cname, mess.Namespace,statsus )
	}
}
func Getdeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	ns := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	DPList,err:=pulick.Deploymnet(clientset,ns)
	if err != nil{
		fmt.Println("namespace err: ",err)
		return
	}
	pulick.Get(c,DPList)
}
func DeleteDeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
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

}
func UpdataDeployment(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	num := c.PostForm("num")
	n,_ := strconv.Atoi(num)
	dnum := int32(n)
	// 扩缩容操作
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
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
}
func GetNode(c *gin.Context)  {
	ENV := c.PostForm("env")
	clientset,_ := clientsetmap[ENV]
	nodes,err:=pulick.Nodes(clientset)
	if err != nil{
		fmt.Println("Nodes err: ",err)
		return
	}
	pulick.Get(c,nodes)
}

func Upload(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	name := c.PostForm("testname")
	fmt.Println("name: ",name)
	file, err := c.FormFile("files")
	if err != nil{
		fmt.Println("err: ",err)
		return
	}
	fmt.Println("file: ",file)
	dst := fmt.Sprintf("%s/%s", Conf.Path, name)
	c.SaveUploadedFile(file, dst)
	Addclientsetmap(name,dst)
	Getnames(c)

}
func Getnames(c *gin.Context)  {
	var names []model.Env
	env := model.Env{}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	filepath.Walk(Conf.Path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Println("path: ",path)
			env.Name = info.Name()
			env.Url = info.Name()
			names = append(names, env)
		}
		return nil
	})
	pulick.Get(c,names)
}
func Delnames(c *gin.Context)  {
	name := c.PostForm("name")
	dst := fmt.Sprintf("%s/%s", Conf.Path, name)
	Delclientsetmap(name,dst)
	Getnames(c)
}
