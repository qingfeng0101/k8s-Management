package k8s

import (
	"context"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func Getdeployment(clientset *kubernetes.Clientset,label,namespace string) ([]appsv1.Deployment,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	 Deploymentlist,err := clientset.AppsV1().Deployments(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return Deploymentlist.Items,nil

}
