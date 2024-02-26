package core

import (
	"fmt"
	"go-blog/server/config"
	"io/ioutil"
	"log"
)

/**
读配置文件
*/

func InitConfig() {
	const configFile = "settings.yaml"
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("get yaml file error: %s", err))
	}

	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("config init unmarshal: %v", err)
	}
	log.Println("config yaml file load init success")
	//global.Config = c
}
