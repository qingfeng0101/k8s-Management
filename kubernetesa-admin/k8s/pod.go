package k8s

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"

)

func GetPod(clientset *kubernetes.Clientset,label,namespace string) ([]corev1.Pod,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	 Podlist,err := clientset.CoreV1().Pods(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return Podlist.Items,nil

}
func GetPodInfo(clientset *kubernetes.Clientset,namespace,podname string) (*corev1.Pod,error) {
	opts := metav1.GetOptions{}

	Podinfo,err := clientset.CoreV1().Pods(namespace).Get(context.Background(),podname,opts)
	if err != nil{
		fmt.Println(err)
		return nil,err
	}
	return Podinfo,nil

}
