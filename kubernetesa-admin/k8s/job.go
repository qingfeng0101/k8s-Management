package k8s

import (
	"context"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetJob(clientset *kubernetes.Clientset,label,namespace string) ([]batchv1.Job,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	Joblist,err := clientset.BatchV1().Jobs(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return Joblist.Items,nil

}
