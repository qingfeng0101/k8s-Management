package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPersistentVolume(clientset *kubernetes.Clientset,label string) ([]corev1.PersistentVolume,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	PersistentVolumelist,err := clientset.CoreV1().PersistentVolumes().List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return PersistentVolumelist.Items,nil

}
