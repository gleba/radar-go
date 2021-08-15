package constants

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

var fileEnv map[string]string

func initfile() {
	fileEnv = make(map[string]string)
	file, err := ioutil.ReadFile("./-env.yaml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(file, &fileEnv)
	if err != nil {
		fmt.Println("ERROR IN Config file : _env.yaml")
		panic(err)
	}

	//vfile, _ := ioutil.ReadFile("./version")
	//Version = string(vfile)
}
func ReadEnv(name string, orValue ...string) string {
	v := os.Getenv(name)
	if v != "" {
		fmt.Println("ENV:", name, "=", v)
		return v
	}
	if fileEnv == nil {
		initfile()
	}
	v = fileEnv[name]
	if v == "" && cap(orValue) > 0 {
		fmt.Println("ENV from file:", name, "=", v)
		v = orValue[0]
	}
	if v != "" {
		fmt.Println("ENV default:", name, "=", v)
		//} else {
		//fmt.Println("ENV:", name, "  default")
	}
	return v
}
