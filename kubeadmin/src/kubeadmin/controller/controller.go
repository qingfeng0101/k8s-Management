package controller

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeadmin/model"
	"net/http"

	"encoding/json"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

func Namespace(c *gin.Context)  {
	var namespacelist []model.NameSapce
	ns := model.NameSapce{}

	nslist,err := clientset.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		fmt.Println("clientset err: ",err)
		return
	}
	for _,NsList := range nslist.Items{
		ns.Name = NsList.Name
		ns.Status = NsList.Status.Phase
        ns.Age = NsList.CreationTimestamp
		namespacelist = append(namespacelist, ns)
	}
	//fmt.Println(namespacelist)
	Get(c,namespacelist)
}
func Pod(c *gin.Context)  {
	// 初始化返回给前端的pod信息的结构体
	var PodList []model.Pods
	pod := model.Pods{}
	// 获取前端传过来的参数
	ns := c.PostForm("namespace")
	//fmt.Println("命名空间： ",ns)
	// 查询pods
	pods := clientset.CoreV1().Pods(ns)
	ps,err := pods.List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return
	}

	for _,p := range ps.Items{

		pod.Name = p.Name
		pod.Namespace = p.Namespace
		pod.Status = p.Status.Phase
		pod.ContainerNum = len(p.Status.ContainerStatuses)
		ReadyNum:=0
		for n:=0;n<len(p.Status.ContainerStatuses);n++{
			pod.RestartNum = p.Status.ContainerStatuses[n].RestartCount
			if p.Status.ContainerStatuses[n].Ready {
				ReadyNum++
			}
		}
		pod.Ready = ReadyNum
		pod.Age = p.CreationTimestamp
		PodList = append(PodList, pod)
	 }
	Get(c,PodList)

}

func DeletePod(c *gin.Context)  {
	var PodList []model.Pods
	pod := model.Pods{}
	// 接受前端删除参数
	c.Request.ParseMultipartForm(128)
	data := c.Request.Form
	name := data["name"][0]
	namespace := data["namespace"][0]
	// 删除操作
	pods := clientset.CoreV1().Pods(namespace)
	err:=pods.Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return
	}
	// 删除后查询pods列表，返回给前端
	pods = clientset.CoreV1().Pods(namespace)
	ps,err := pods.List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return
	}
	//
	for _,p := range ps.Items{
		fmt.Println(p.Name)
		fmt.Println(p.Status.Phase,p.OwnerReferences)

		pod.Name = p.Name
		pod.Namespace = p.Namespace
		pod.Status = p.Status.Phase
		pod.ContainerNum = len(p.Status.ContainerStatuses)
		ReadyNum:=0
		for n:=0;n<len(p.Status.ContainerStatuses);n++{
			pod.RestartNum = p.Status.ContainerStatuses[n].RestartCount
			if p.Status.ContainerStatuses[n].Ready {
				ReadyNum++
			}
		}
		pod.Ready = ReadyNum
		pod.Age = p.CreationTimestamp
		PodList = append(PodList, pod)
	}
	Get(c,PodList)

}
func GetPodInfo(c *gin.Context)  {

	// 实例化pod
	pod := model.Pods{}
	// 接受前端删除参数
	c.Request.ParseMultipartForm(128)
	data := c.Request.Form
	name := data["name"][0]
	namespace := data["namespace"][0]

	// 获取对应pod的信息，并且复制给实例化的pod
	pods := clientset.CoreV1().Pods(namespace)
	ps,err:=pods.Get(context.TODO(),name,metav1.GetOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return
	}
	pod.Name = ps.Name
	pod.Namespace = ps.Namespace
	pod.Labels = ps.Labels
	pod.Age = ps.CreationTimestamp
	if len(ps.Status.Message) == 0{
		pod.Message = "none"
	} else {
		pod.Message = ps.Status.Message
	}
	if len(ps.Status.Reason) == 0{
		pod.Reason = "none"
	} else {
		pod.Reason = ps.Status.Reason
	}
	pod.Status = ps.Status.Phase
	pod.HostIP = ps.Status.HostIP
	pod.PodIP = ps.Status.PodIP
	pod.QOSClass = ps.Status.QOSClass
	// 因为多个容器是一个数组所以这里用遍历的方式
	fmt.Println()
	for n:=0;n<len(ps.Status.ContainerStatuses);n++{
		container := model.Container{}
		container.Ready = ps.Status.ContainerStatuses[n].Ready
		container.ContainerID = ps.Status.ContainerStatuses[n].ContainerID
		container.Name = ps.Status.ContainerStatuses[n].Name
		container.Image = ps.Status.ContainerStatuses[n].Image
		container.ImageID = ps.Status.ContainerStatuses[n].ImageID
		container.Status = ps.Status.ContainerStatuses[n].State
		container.RestartNum = ps.Status.ContainerStatuses[n].RestartCount
		container.Ports =  ps.Spec.Containers[n].Ports
		container.Lcpu = ps.Spec.Containers[n].Resources.Limits.Cpu()
		container.Lmem = ps.Spec.Containers[n].Resources.Limits.Memory()
		container.Rcpu = ps.Spec.Containers[n].Resources.Requests.Cpu()
		container.Rmem = ps.Spec.Containers[n].Resources.Requests.Memory()
		pod.Containers = append(pod.Containers,container )
	}
    // 将赋值完的pod实例返回给前端
	Get(c,pod)
}
func GetPodlog(c *gin.Context)  {

	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	var lines int64
	lines = 100

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
		pods := clientset.CoreV1().Pods(mess.Namespace)
		logs :=pods.GetLogs(mess.Name,&v1.PodLogOptions{Follow: true,TailLines: &lines})
		log,err := logs.Stream(context.TODO())
		defer log.Close()
		if err != nil{
			fmt.Println(err)
			return
		}
		r := bufio.NewReader(log)
		for {
			bytes, err := r.ReadBytes('\n')

            fmt.Println(string(bytes))
			err = ws.WriteMessage(mt, bytes)
			if err != nil {
				break
			}

			//Get(c,string(bytes))
			if err != nil {
				if err != io.EOF {
					fmt.Println("err: ", err)
					return
				}
				return
			}
		}
	}



}