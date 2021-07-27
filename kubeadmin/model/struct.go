package model

import (
	"fmt"
	"io/ioutil"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"gopkg.in/yaml.v2"
)


type NameSapce struct {
	Name string `json:"name"`
	Status interface{} `json:"status"`
	Age v1.Time `json:"age"`
}
type Pods struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Labels map[string]string `json:"labels"`
	Message string `json:"message"`
	Reason string `json:"reason"`
	HostIP string `json:"host_ip"`
	PodIP string `json:"pod_ip"`
	QOSClass interface{} `json:"qos_class"`
	Age v1.Time `json:"age"`
	Ready int `json:"ready"`
	ContainerNum int `json:"container_num"`
	RestartNum int32 `json:"restart_num"`
	Status interface{} `json:"status"`
	Containers []Container `json:"containers"`
	ReadyStr string `json:"ready_str"`

}
type Container struct {
	Name string
	Image string
	ImageID string
	ContainerID string
	Ready  bool
	Status interface{}
	RestartNum int32
	Ports interface{}
	Lcpu interface{}
	Lmem interface{}
	Rcpu interface{}
	Rmem interface{}
}
type Message struct {
	Name  string `json:"name"`
	Namespace string `json:"namespace"`
	Cname string `json:"cname"`
	ENV string `json:"env"`
	Status string `json:"status"`
}
type Deployment struct {
	NAME string `json:"name"`
	NAMESPACE string `json:"namespace"`
	DESIRED int32 `json:"desired"`
	CURRENT int32 `json:"current"`
	UPTODATE int32 `json:"uptodate"`
	AVAILABLE int32 `json:"available"`
	AGE v1.Time `json:"age"`
	ISSHOW bool `json:"isshow"`
}
type Node struct {
	NAME string `json:"name"`
	STATUS interface{} `json:"status"`
	ROLES string `json:"roles"`
	AGE v1.Time `json:"age"`
	VERSION string `json:"version"`
}
type Env struct {
	Name string `json:"name"`
	Url string `json:"url"`
}
type BaseInfo struct {
	Ip string `yaml:"ip"`
	Port string `yaml:"port"`
	Path    string `yaml:"path"`

}
func (c *BaseInfo) GetConf(path string) *BaseInfo {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		fmt.Println(err.Error())
	}
	return c
}