package k8s

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetPersistentVolumeClaim(clientset *kubernetes.Clientset,label,namespace string) ([]corev1.PersistentVolumeClaim,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	PersistentVolumeClaimlist,err := clientset.CoreV1().PersistentVolumeClaims(namespace).List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return PersistentVolumeClaimlist.Items,nil

}
