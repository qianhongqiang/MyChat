package conf

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/qianhongqiang/MyChat/common/conf"
)

var (
	confPath string
	Conf *Config
)

type Config struct {
	*conf.CommonConf
	confFile string
	Server *conf.Server
	ServiceDiscoveryServer *conf.ServiceDiscoveryServer
}

func init()  {
	//flag.StringVar(&confPath,"conf","./access.toml","config path")
	flag.StringVar(&confPath,"conf","./server/access/access.toml","config path")
}

func Init() (err error) {
	_,err = toml.DecodeFile(confPath,&Conf)
	return
}