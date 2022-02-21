package xsls

import (
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"time"
)

var producerInstance *producer.Producer

func SetProducerInstance(producer *producer.Producer) {
	producerInstance = producer
}

//写入阿里云日志
func AliyunSendLog(yunLogProject, yunLogStore, sourceIP, addLog string) error {

	producerInstance.Start()
	log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": addLog})
	//sourceIP,_:=util.ExternalIP()
	err := producerInstance.SendLog(yunLogProject, yunLogStore, "topic", sourceIP, log)
	if err != nil {
		return err
	}
	//err1:=producerInstance.Close(60)
	return nil
}
func SetProducerClose() {
	producerInstance.SafeClose()
}
