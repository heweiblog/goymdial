package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

func main() {
	cfg, err := goconfig.LoadConfigFile("/home/heweiwei/go/src/goymdial/conf/goymdial.ini")
	if err != nil {
		panic("错误")
	}
	value, err := cfg.GetValue("dial", "server")
	if err != nil {
		panic("错误")
	} else {
		fmt.Println(value)
	}

	valut, err := cfg.Int("agent", "port")
	if err != nil {
		panic("错误")
	} else {
		fmt.Println(valut)
	}

}
