package server

import (
	"github.com/qianhongqiang/MyChat/libnet"
	"github.com/coreos/etcd/etcdserver"
	"github.com/oikomi/FishChatServer2/conf_discovery/etcd"
)

type Server struct {
	Server *libnet.Server
}

func New() (s *Server) {
	s = &Server{}
	return
}

func SDHeart()  {

}