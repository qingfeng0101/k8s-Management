package global

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	"os"
	"path/filepath"
)


var clientset *kubernetes.Clientset
var metricsset  *metrics.Clientset
//var configclient  *rest.Config
var configMap map[string]*kubernetes.Clientset
var metricsMap map[string]*metrics.Clientset
var confMap map[string]*rest.Config
var config *rest.Config

func initClient(path string) error {
	//var err error

	configMap = make(map[string]*kubernetes.Clientset)
	metricsMap = make(map[string]*metrics.Clientset)
	confMap = make(map[string]*rest.Config)
	//var ClusterNames []string
	//kubeconfig := filepath.Join(homedir.HomeDir(),".kube","config")
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			//fmt.Println("path: ",path)
			//env.Name = info.Name()
			//env.Url = info.Name()
			//names = append(names, env)
			//fmt.Println("name: ", info.Name())
			//ClusterNames = append(ClusterNames, info.Name())
			if config, err = clientcmd.BuildConfigFromFlags("",path);err != nil{
				return err
			}
			if clientset,err = kubernetes.NewForConfig(config);err != nil{
				return err
			}
			if metricsset,err = metrics.NewForConfig(config);err != nil{
				return err
			}
			configMap[info.Name()] = clientset
			metricsMap[info.Name()] = metricsset
			confMap[info.Name()] = config
		}


		return nil
	})
	//if config, err = rest.InClusterConfig();err != nil{
	//	if config, err = clientcmd.BuildConfigFromFlags("",kubeconfig);err != nil{
	//		return err
	//	}
	//}
	//if clientset,err = kubernetes.NewForConfig(config);err != nil{
	//	return err
	//}
	//if metricsset,err = metrics.NewForConfig(config);err != nil{
	//	return err
	//}
	return nil
}
func K8sclientset(env string)  (*kubernetes.Clientset) {
	clientset = configMap[env]
	return clientset
}
func K8sconf(env string)  (*rest.Config) {
	config = confMap[env]
	return config
}
func K8smetricsset(env string)  ( *metrics.Clientset) {
	metricsset = metricsMap[env]
	return metricsset
}