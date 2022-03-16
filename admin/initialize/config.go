package initialize

import (
	"catering/global"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func InitConfig() {
	path := "./conf/config.yaml"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("err", err)
	}

	err = yaml.Unmarshal(data, &global.Config)
	if err != nil {
		fmt.Println("err", err)
	}
}
