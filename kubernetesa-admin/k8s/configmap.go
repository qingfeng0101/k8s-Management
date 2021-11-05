package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetConfigMap(clientset *kubernetes.Clientset,label,namespace string) ([]corev1.ConfigMap,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	ConfigMaplist,err := clientset.CoreV1().ConfigMaps(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return ConfigMaplist.Items,nil

}
