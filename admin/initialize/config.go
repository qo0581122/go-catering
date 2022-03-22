package initialize

import (
	"catering/global"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func InitConfig() error {
	path := "./conf/config.yaml"
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &global.Config)
	if err != nil {
		return err
	}
	return nil
}
