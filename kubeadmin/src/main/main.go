package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)
type BaseInfo struct {
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
func main() {
	var path string
	flag.StringVar(&path,"path","/usr/down","upload config dir")
	flag.Parse()
	fmt.Println(path)
	info := BaseInfo{}
	conf := info.GetConf(path)
	fmt.Println("path: ",conf.Path)
}
