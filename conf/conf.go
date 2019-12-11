package conf

import (
	"github.com/Unknwon/goconfig"
)

type Config struct {
	DialServer  string
	AgentPort   int
	AgentIp     string
	Health      int
	DelayWeight int
	LostWeight  int
	Count       int
	Timeout     int
	Interval    int
	Dname       string
}

var Conf *Config

func init() {
	cfg, err := goconfig.LoadConfigFile("/home/heweiwei/go/src/goymdial/conf/goymdial.ini")
	if err != nil {
		panic("load conf error")
	}

	Conf = new(Config)
	Conf.DialServer, err = cfg.GetValue("dial", "server")
	if err != nil {
		panic("load conf error")
	}

	Conf.AgentPort, err = cfg.Int("agent", "port")
	if err != nil {
		panic("load conf error")
	}
}
