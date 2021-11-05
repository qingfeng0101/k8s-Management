package k8s

import (
	"context"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)
var resourece map[string]int64
var memnum map[string]string
func GetNode(clientset *kubernetes.Clientset,metricsset *metrics.Clientset,label string) ([]core.Node, map[string]int64,map[string]string,error) {
     opts := metav1.ListOptions{}
     if label != ""{
     	opts.LabelSelector = label
	 }
	 nodelist,err := clientset.CoreV1().Nodes().List(context.Background(),opts)
	 if err != nil{
		 return nil,nil,nil,err
	 }
	nodelistmetrics,err := metricsset.MetricsV1beta1().NodeMetricses().List(context.Background(),opts)
	if err != nil{
		return nil,nil,nil,err
	}
	//nodemetrics,err := metricsset.MetricsV1beta1().NodeMetricses().Get(context.Background(),"master-15-01",metav1.GetOptions{})
	//if err != nil{
	//	return nil,nil,nil,err
	//}

	//cpu := nodemetrics.Usage.Cpu().MilliValue()
	resourece = make(map[string]int64)
	memnum = make(map[string]string)
	for n :=0; n<len(nodelistmetrics.Items);n++{
		cpunum := nodelistmetrics.Items[n].Usage.Cpu().MilliValue()
		mem := nodelistmetrics.Items[n].Usage.Memory().String()
		resourece[nodelistmetrics.Items[n].Name] = cpunum
		memnum[nodelistmetrics.Items[n].Name]=mem
	}

	//fmt.Println("nodemetrics: ",nodemetrics.Usage.Memory())
	//fmt.Println("nodemetrics: ",memnum)
	//fmt.Println("nodemetrics: ",nodemetrics.Usage.Cpu(),cpu,resourece)
	return nodelist.Items,resourece,memnum,nil

}
