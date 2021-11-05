package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNamespace(clientset *kubernetes.Clientset,label string) ([]corev1.Namespace,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	 Namespacelist,err := clientset.CoreV1().Namespaces().List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return Namespacelist.Items,nil

}
