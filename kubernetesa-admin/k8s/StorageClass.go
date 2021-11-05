package k8s

import (
	"context"

	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetStorageClass(clientset *kubernetes.Clientset,label  string) ([]storagev1.StorageClass,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	StorageClasslist,err := clientset.StorageV1().StorageClasses().List(context.Background(),opts)
	 if err != nil{
		 return nil,err
	 }
	return StorageClasslist.Items,nil

}
