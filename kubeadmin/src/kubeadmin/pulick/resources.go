package pulick

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"kubeadmin/model"
	"strconv"
)

func NamespaceList(clientset *kubernetes.Clientset) ([]model.NameSapce,error) {
	var namespacelist []model.NameSapce
	ns := model.NameSapce{}
	nslist,err := clientset.CoreV1().Namespaces().List(context.TODO(),metav1.ListOptions{})
		if err != nil{
			fmt.Println("clientset err: ",err)
			return nil,err
		}
		for _,NsList := range nslist.Items{
			ns.Name = NsList.Name
			ns.Status = NsList.Status.Phase
			ns.Age = NsList.CreationTimestamp
			namespacelist = append(namespacelist, ns)
		}
		return namespacelist,nil
}
func Pods(clientset *kubernetes.Clientset,namespace string) ([]model.Pods,error) {
	// 初始化返回给前端的pod信息的结构体
	var PodList []model.Pods
	pod := model.Pods{}
	// 查询pods
	pods := clientset.CoreV1().Pods(namespace)
	ps,err := pods.List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		return nil,err
	}
	for _,p := range ps.Items{
		pod.Name = p.Name
		pod.Namespace = p.Namespace
		pod.ContainerNum = len(p.Status.ContainerStatuses)
		ReadyNum:=0
		for n:=0;n<len(p.Status.ContainerStatuses);n++{
			pod.RestartNum = p.Status.ContainerStatuses[n].RestartCount
			var con v1.ContainerStateWaiting
			j,err :=json.Marshal(p.Status.ContainerStatuses[0].State.Waiting)
			if err != nil {
				return nil,err
			}
			if string(j) != "null"{
				//fmt.Println("j: ",string(j))
				error:=json.Unmarshal(j,&con)
				if error != nil{
					fmt.Println("Unmarshal err: ",error)
					return nil,err
				}
				pod.Status = con.Reason
			}else {
				pod.Status = p.Status.Phase
			}

			//fmt.Println(p.Status.ContainerStatuses[0].State.Waiting)
			//fmt.Println(p.Status.Conditions[0].Reason)
			if p.Status.ContainerStatuses[n].Ready {
				ReadyNum++
			}
		}
		pod.Ready = ReadyNum
		pod.Age = p.CreationTimestamp
		pod.ReadyStr = strconv.Itoa(ReadyNum) + "/" + strconv.Itoa(pod.ContainerNum)
		PodList = append(PodList, pod)
	}
	return PodList,nil
}
func Deploymnet(clientset *kubernetes.Clientset,namespace string) ([]model.Deployment,error) {
	// 初始化返回给前端的pod信息的结构体
	var DPList []model.Deployment
	DP := model.Deployment{}
	// 获取前端传过来的参数

	//fmt.Println("命名空间： ",ns)
	// 查询pods
	dps := clientset.AppsV1().Deployments(namespace)
	ds,err := dps.List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		return nil,err
	}

	for _,d := range ds.Items{
		DP.NAME = d.Name
		DP.NAMESPACE = d.Namespace
		DP.AGE = d.CreationTimestamp
		DP.AVAILABLE = d.Status.AvailableReplicas
		DP.CURRENT = d.Status.Replicas
		DP.DESIRED = d.Status.Replicas
		DP.UPTODATE = d.Status.UpdatedReplicas
		DPList = append(DPList, DP)
	}
	return DPList,nil
}
func Nodes(clientset *kubernetes.Clientset) ([]model.Node,error) {
	// 初始化返回给前端的pod信息的结构体
	var NodeList []model.Node
	Node := model.Node{}

	// 查询pods
	Nodes := clientset.CoreV1().Nodes()
	Ns,err := Nodes.List(context.TODO(),metav1.ListOptions{})
	if err != nil{
		return nil,err
	}

	for _,n := range Ns.Items{
		Node.NAME = n.Name
		Node.STATUS = n.Status.Conditions[3].Type
		Node.AGE = n.CreationTimestamp
		Node.VERSION = n.Status.NodeInfo.KubeletVersion
		NodeList = append(NodeList, Node)
	}
	return NodeList,nil
}
func PodLogs(ws *websocket.Conn,mt int,clientset *kubernetes.Clientset,name,namespace string)  {
	var lines int64
	lines = 100
	pods := clientset.CoreV1().Pods(namespace)
	logs :=pods.GetLogs(name,&v1.PodLogOptions{Follow: true,TailLines: &lines})
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
func PodINFO(clientset *kubernetes.Clientset,name,namespace string) (model.Pods,error) {
	// 实例化pod
	pod := model.Pods{}
	// 获取对应pod的信息，并且复制给实例化的pod
	pods := clientset.CoreV1().Pods(namespace)
	ps,err:=pods.Get(context.TODO(),name,metav1.GetOptions{})
	if err != nil{
		return pod,err
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
	return pod,nil

}
func DeletePod(clientset *kubernetes.Clientset,name,namespace string) (bool,error) {
	// 删除操作
	pods := clientset.CoreV1().Pods(namespace)
	err:=pods.Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return false,err
	}
	return true,nil
}
func DeleteDeployment(clientset *kubernetes.Clientset,name,namespace string) (bool,error) {

	// 删除操作
	dps := clientset.AppsV1().Deployments(namespace)
	err:=dps.Delete(context.TODO(),name,metav1.DeleteOptions{})
	if err != nil{
		fmt.Println("pods err: ",err)
		return false,err
	}
	return true,nil
}
func UpDeployment(clientset *kubernetes.Clientset,name,namespace string,num *int32) (bool,error) {
	dps := clientset.AppsV1().Deployments(namespace)
	dp,err:=dps.Get(context.TODO(),name,metav1.GetOptions{})
	if err != nil{
		return false,err
	}
	dp.Spec.Replicas = num
	_,err =dps.Update(context.TODO(),dp,metav1.UpdateOptions{})
	if err != nil{
		return false,err
	}
	return true,nil
}