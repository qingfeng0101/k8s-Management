package model

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


type NameSapce struct {
	Name string
	Status interface{}
	Age v1.Time
}
type Pods struct {
	Name string
	Namespace string
	Labels map[string]string
	Message string
	Reason string
	HostIP string
	PodIP string
	QOSClass interface{}
	Age v1.Time
	Ready int
	ContainerNum int
	RestartNum int32
	Status interface{}
	Containers []Container

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
}
