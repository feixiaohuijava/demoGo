package main

import (
	"demoGo/toolConfig"
	"encoding/json"
	"fmt"
	"path"
	"runtime"
)

func getCurrentPath() string {
	_, filename, _, ok := runtime.Caller(1)
	var currentPath string
	if ok {
		currentPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		currentPath = "./"
	}
	return currentPath
}

func main() {

	currentPath := getCurrentPath()
	var c toolConfig.Conf
	//读取yaml配置文件
	conf := c.GetConf(currentPath)
	fmt.Println(conf)

	//将对象，转换成json格式
	data, err := json.Marshal(conf)

	if err != nil {
		fmt.Println("err:\t", err.Error())
		return
	}

	//最终以json格式，输出
	fmt.Println("data:\t", string(data))
}
