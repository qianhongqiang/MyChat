package conf

import "github.com/oikomi/FishChatServer2/common/xtime"

type CommonConf struct {
	Ver string
	LogPath string
}

type Server struct {
	Proto string
	Addr string
}

type ServiceDiscoveryServer struct {
	ServiceName string
	RPCAddr     string
	EtcdAddr    string
	Interval    xtime.Duration
	TTL         xtime.Duration
}

type Zookeeper struct {
	Root    string
	Addrs   []string
	Timeout xtime.Duration
}

type KafkaProducer struct {
	Zookeeper *Zookeeper
	Brokers   []string
	Sync      bool // true: sync, false: async
}