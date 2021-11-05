package controller

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"io"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"kubernetesa-admin/global"
	"net/http"
)
var upGrader = websocket.Upgrader{
	CheckOrigin: func (r *http.Request) bool {
		return true
	},
}

type Message struct {
	Name  string `json:"name"`
	Namespace string `json:"namespace"`
	Cname string `json:"cname"`
	ENV string `json:"env"`

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
		var mess  Message
		//var d map[string]interface{}
		// var i interface{}
		fmt.Println("message",message)
		err = json.Unmarshal(message, &mess)
		if err != nil{
			fmt.Println("Unmarshal err: ",err)
			return
		}


		clientset := global.K8sclientset(mess.ENV)

		go PodLogs(ws, mt, clientset, mess.Name,mess.Cname, mess.Namespace,statsus )
	}
}
func PodLogs(ws *websocket.Conn,mt int,clientset *kubernetes.Clientset,name,cname,namespace string,status chan bool)  {
	var lines int64
	lines = 100
	pods := clientset.CoreV1().Pods(namespace)
	logs :=pods.GetLogs(name,&v1.PodLogOptions{Follow: true,Container: cname,TailLines: &lines})
	defer ws.Close()
	fmt.Println("cname: ",cname)
	fmt.Println("name: ",name)
	log,err := logs.Stream(context.TODO())
	defer log.Close()
	if err != nil{
		fmt.Println(err)
		return
	}
	r := bufio.NewReader(log)
	for {

		select {

		case <- status:

			return
		default:
			//fmt.Println("第二层",mess.Status)

			bytes, err := r.ReadBytes('\n')
			//fmt.Println("第二层 ",status)
			//fmt.Println(string(bytes))
			//fmt.Println("********************************************************************")
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