package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
	"kubernetesa-admin/global"
	"kubernetesa-admin/global/terminal"

	wsterminal "kubernetesa-admin/global/websocket"
	"log"
)




func PodExec(c *gin.Context)  {
	cmd  := []string{"/bin/sh"}
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ws.Close()
	_, message, err := ws.ReadMessage()
	var mess  Message
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
	clientset := global.K8sclientset(mess.ENV)

	config := global.K8sconf(mess.ENV)
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
	err = executor.Stream(remotecommand.StreamOptions{
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
