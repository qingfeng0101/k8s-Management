package k8s

import (
	"context"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetCrontabJob(clientset *kubernetes.Clientset,label,namespace string) ([]batchv1.CronJob,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	CrontabJoblist,err := clientset.BatchV1().CronJobs(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return CrontabJoblist.Items,nil

}
