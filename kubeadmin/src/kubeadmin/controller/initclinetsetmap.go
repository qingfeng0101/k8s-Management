package controller

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func initclientsetmap(path string)  {
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			info.Name()
			k8sconfig := flag.String(info.Name(),path,"kubernetes config file path")
			flag.Parse()

			config , err := clientcmd.BuildConfigFromFlags("",*k8sconfig)
			if err != nil {
				log.Println(err)
			}
			clientset , err = kubernetes.NewForConfig(config)
			if err != nil {
				log.Fatalln(err)
			} else {
				clientsetmap[info.Name()]=clientset
				configmap[info.Name()]=config
				fmt.Println("connect k8s success")
			}

		}
		fmt.Println(len(clientsetmap))
		return nil
	})
}
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}
func Addclientsetmap(name,path string)  {
	_,OK := clientsetmap[name]
	if !OK {
		fs := flag.NewFlagSet(name,flag.ExitOnError)
		var k8sconfig *string
		k8sconfig = fs.String(name,path,"kubernetes config file path")
		fmt.Println("k8sconfig1",k8sconfig)
		flag.Parse()
        fmt.Println("k8sconfig2",k8sconfig)
		config , err := clientcmd.BuildConfigFromFlags("",*k8sconfig)
		if err != nil {
			log.Println(err)
		}
		clientset , err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Fatalln(err)
		} else {
			clientsetmap[name]=clientset
			fmt.Printf("connect %s k8s success\n",name)
		}
	}else {
		fmt.Println("当前环境存在")
	}
}
func Delclientsetmap(name,path string)  {
	err := os.Remove(path)
	if err != nil{
		fmt.Println("Remove err: ",err)
		return
	}else {
		delete(clientsetmap,name)
	}
	fmt.Println("当前环境数： ",len(clientsetmap))
}
