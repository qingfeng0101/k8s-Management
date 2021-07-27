package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/terminal"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"kubeadmin/model"
	"kubeadmin/pulick"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
wsterminal "github.com/websocket"
	"context"

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
func GetDeploymentInfo(c *gin.Context)  {
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	Dpinfo,err:=pulick.DeploymentINFO(clientset,name,namespace)
	if err != nil{
		fmt.Println("PodINFO err: ",err)
		return
	}
	pulick.Get(c,Dpinfo)
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
func UpdataImages(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")
	ENV := c.PostForm("env")
	name := c.PostForm("name")
	namespace := c.PostForm("namespace")
	num := c.PostForm("num")
	images := c.PostForm("images")
	n,_ := strconv.Atoi(num)
	dnum := int32(n)
	// 扩缩容操作
	clientset,OK := clientsetmap[ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	OK,err:=pulick.UpPod(clientset,name,namespace,images,dnum)
	if err != nil{
		fmt.Println("DeletePod err: ",err)
		return
	}
	if OK{
		// 扩缩容成功后查询返回给前端
		podinfo,err := pulick.PodINFO(clientset,name,namespace)
		if err != nil{
			fmt.Println("Pods err: ",err)
			return
		}
		pulick.Get(c,podinfo)
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

type named struct {

}
func EXEC(c *gin.Context)  {
	cmd  := []string{"/bin/sh"}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	_, message, err := ws.ReadMessage()
	var mess  model.Message
	err = json.Unmarshal(message, &mess)
	if err != nil{
		fmt.Println("Unmarshal err: ",err)
		return
	}
	namespace := mess.Namespace
	podName := mess.Name
	containerName := mess.Cname
	//ctx, stop := context.WithCancel(context.Background())
	pty, err := wsterminal.NewTerminalSession(ws)
	if err != nil {
		log.Printf("get pty failed: %v\n", err)
		return
	}
	defer func() {
		log.Println("close session.")
		pty.Close()
	}()
	clientset,OK := clientsetmap[mess.ENV]
	if !OK{
		fmt.Println("当前环境不存在")
		return
	}
	config,_ := configmap[mess.ENV]
	pod, err := PodGet(clientset,podName, namespace)
	if err != nil {
		log.Printf("get kubernetes client failed: %v\n", err)
		return
	}
	ok, err := terminal.ValidatePod(pod, containerName)
	if !ok {
		msg := fmt.Sprintf("Validate pod error! err: %v", err)
		log.Println(msg)
		pty.Write([]byte(msg))
		pty.Done()
		return
	}
	err = Podexec(clientset,config,cmd, pty, namespace, podName, containerName)
	if err != nil {
		msg := fmt.Sprintf("Exec to pod error! err: %v", err)
		log.Println(msg)
		pty.Write([]byte(msg))
		pty.Done()
	}
	return
}
func Podexec(clientset *kubernetes.Clientset,config *rest.Config,cmd []string, ptyHandler terminal.PtyHandler, namespace, podName, containerName string) error  {
	defer func() {
		ptyHandler.Done()
	}()

	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec")

	req.VersionedParams(&corev1.PodExecOptions{
		Container: containerName,
		Command:   cmd,
		Stdin:     !(ptyHandler.Stdin() == nil),
		Stdout:    !(ptyHandler.Stdout() == nil),
		Stderr:    !(ptyHandler.Stderr() == nil),
		TTY:       ptyHandler.Tty(),
	}, scheme.ParameterCodec)

	executor, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return err
	}
	err = executor.Strlseam(remotecommand.StreamOptions{
		Stdin:             ptyHandler.Stdin(),
		Stdout:            ptyHandler.Stdout(),
		Stderr:            ptyHandler.Stderr(),
		TerminalSizeQueue: ptyHandler,
		Tty:               ptyHandler.Tty(),
	})
	return err
}
func  PodGet(clientset *kubernetes.Clientset,name, namespace string) (*corev1.Pod, error) {
	opt := metav1.GetOptions{}
	return clientset.CoreV1().Pods(namespace).Get(context.TODO(),name, opt)
}


