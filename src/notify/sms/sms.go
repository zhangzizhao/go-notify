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
	//        fmt.Println(innerSend("10.10.10.87:9090",message))
	fmt.Println(innerSend("10.231.144.37:9090", message))
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
