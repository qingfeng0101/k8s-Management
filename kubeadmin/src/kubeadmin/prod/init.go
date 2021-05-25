package prod

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

var clientset *kubernetes.Clientset

func init()  {

	prodconfig := flag.String("prodconfig","/root/pord/config","kubernetes config file path")
	flag.Parse()

	config , err := clientcmd.BuildConfigFromFlags("",*prodconfig)
	if err != nil {
		log.Println(err)
	}
	clientset , err = kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("connect k8s success")
	}

}
