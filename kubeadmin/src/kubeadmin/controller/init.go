package controller

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)
// 测试环境的clientset
var clientset *kubernetes.Clientset
// 生产环境的prodclientset
var prodclientset *kubernetes.Clientset

func init()  {
    // 通过读取测试的config文件生成clientset
	k8sconfig := flag.String("k8sconfig","/Users/zhouhao/kubeadmin/src/kubeadmin/controller/config","kubernetes config file path")
	flag.Parse()

	config , err := clientcmd.BuildConfigFromFlags("",*k8sconfig)
	if err != nil {
		log.Println(err)
	}
	clientset , err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connect k8s success")
	}
	// 通过读取线上的config文件生成clientset
	prodconfig := flag.String("prodconfig","/Users/zhouhao/kubeadmin/src/kubeadmin/prod/config","kubernetes config file path")
	flag.Parse()

	pconfig , err := clientcmd.BuildConfigFromFlags("",*prodconfig)
	if err != nil {
		log.Println(err)
	}
	prodclientset , err = kubernetes.NewForConfig(pconfig)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connect k8s success")
	}
}
