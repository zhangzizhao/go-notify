package main

import (
	"fmt"
	"sort"
	"strings"

	"notify/lib"
	"notify/mail"
)

type NotifyType struct {
	Channel string `json:"channel"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Recieve string `json:"receive"`
}

func (c *NotifyType) SendMail(notify lib.NotifyContent) {
	fmt.Println("sand mail:", notify)
	mail.Send(notify)
}

func (c *NotifyType) SendSms(notify lib.NotifyContent) {
	fmt.Println("sand sms:", notify)
}

func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 1 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func DoNotify(notify NotifyType) error {
	var notifyContent lib.NotifyContent
	notifyContent.Subject = notify.Subject
	notifyContent.Content = notify.Content
	notifyContent.Recieve = notify.Recieve
	channelList := strings.Split(notify.Channel, ";")
	sort.Strings(channelList)
	notifyChannel := RemoveDuplicatesAndEmpty(channelList)
	for _, channel := range notifyChannel {
		switch channel {
		case "mail":
			notify.SendMail(notifyContent)
		case "sms":
			notify.SendSms(notifyContent)
		case "hipchat":
			fmt.Println("notify hipchat...  skip")
		default:
			fmt.Println("unknown notify channel and skip:", channel)
		}

	}
    return nil
}
