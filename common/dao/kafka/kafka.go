package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/qianhongqiang/MyChat/common/conf"
	"fmt"
	"time"
)

type Producer struct {
	sarama.AsyncProducer
	sarama.SyncProducer
	c *conf.KafkaProducer
	env string
}

func NewProducer(c *conf.KafkaProducer) (p *Producer)  {
	p = &Producer{
		c:c,
		env:fmt.Sprintf("zookeeper%s@%v|borker$v|sync(%t)",c.Zookeeper.Root,c.Zookeeper.Addrs,c.Brokers,c.Sync),
	}

	if !c.Sync {

	}

	return
}

func (p *Producer) asyncDial() (err error)  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Compression = sarama.CompressionSnappy
	config.Producer.Flush.Frequency = 500 * time.Millisecond
	if p.AsyncProducer, err = sarama.NewAsyncProducer(p.c.Brokers,config); err == nil {

	}
	return
}