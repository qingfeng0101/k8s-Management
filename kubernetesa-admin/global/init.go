package global

import (
	"flag"
	"fmt"
	"io/ioutil"
	"kubernetesa-admin/global/conf"
)

var Conf *conf.BaseInfo
var status = false
func Init()  error {
	if (!status){
		initsystem()
	}
	s, _ := ioutil.ReadDir(Conf.Path)
	fmt.Println("11: ",len(s))
	if len(s) == 0 {
		fmt.Println("empty")
	}else {
		if err := initClient(Conf.Path);err != nil{
			return  err
		}
	}
	return nil
}
func initsystem() bool {
	var path string
	flag.StringVar(&path,"f","/usr/down","config")
	flag.Parse()
	info := conf.BaseInfo{}
	Conf = info.GetConf(path)
	status = true
	return status
}