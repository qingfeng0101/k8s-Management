package conf

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

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

