package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func Reload(filePath string,assignment interface{}) (err error) {

	if filePath == "" {
		return errors.New("配置文件路径为空")
	}
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
	}

	fileData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return
	}

	//将数据解析到struct中
	err = yaml.Unmarshal(fileData, assignment)
	if err != nil {
		return
	}

	return
}
