package k8s

import (
	"context"
	extensionsv1beta1 "k8s.io/api/extensions/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetIngress(clientset *kubernetes.Clientset,label,namespace string) ([]extensionsv1beta1.Ingress,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	Ingresslist,err := clientset.ExtensionsV1beta1().Ingresses(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return Ingresslist.Items,nil

}
