package main

import (
	"flag"
	"github.com/golang/glog"
	"github.com/qianhongqiang/MyChat/server/access/conf"
	"github.com/qianhongqiang/MyChat/server/access/server"
	"github.com/qianhongqiang/MyChat/libnet"
	"github.com/qianhongqiang/MyChat/codec"
	"fmt"
)

type Config struct {
	configFile string
}

func init()  {
	flag.Set("alsologtostderr", "true")
}

func main()  {
	var err error
	flag.Parse()
	if err := conf.Init(); err != nil {
		glog.Error("conf.Init() error: ", err)
		panic(err)
	}
	accessServer := server.New()
	protobuf := codec.Protobuf()
	fmt.Println(conf.Conf.Server.Proto)
	fmt.Println(conf.Conf.Server.Addr)
	if accessServer.Server, err = libnet.Serve(conf.Conf.Server.Proto, conf.Conf.Server.Addr, protobuf, 0); err != nil {
		glog.Error(err)
		panic(err)
	}

}
