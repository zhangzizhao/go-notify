package sms

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"xiaoju/sms"

	"notify/lib"
)

func Send(notify lib.NotifyContent) error {
	message := sms.NewMessage()
	message.BusinessId = 200100000
	//message.Phones = []int64{13521022145, 15600389089}
	message.Phones = notify.PhoneList
	message.Message = notify.Content
	message.JobId = 0
	//        fmt.Println(innerSend("xx.xx.xx.xx:xxxx",message))
	fmt.Println(innerSend("xx.xx.xx.xx:xxxx", message))
	return nil
}

func innerSend(addr string, message *sms.Message) (r sms.SerialNumber, se *sms.ServiceError, ie *sms.InputError, err error) {
	trans, err := thrift.NewTSocket(addr)
	if err != nil {
		return
	}
	defer trans.Close()

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	client := sms.NewMessageServiceClientFactory(trans, protocolFactory)

	if err = trans.Open(); err != nil {
		return
	}

	return client.SendMessage(message)
}
