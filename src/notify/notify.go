package main

import (
	"fmt"
	"sort"
	"strings"

	"notify/lib"
	"notify/mail"
	"notify/sms"
	"notify/hipchat"
)

type NotifyType struct {
	Channel   string  `json:"channel"`
	Subject   string  `json:"subject"`
	Content   string  `json:"content"`
	Recieve   string  `json:"receive"`
	PhoneList []int64 `json:"phonelist"`  // ???  todo
}

func (c *NotifyType) CheckParam() (int64, string) {
	if c.Recieve == "" || c.Channel == "" {
		return 404, "channel or recieve empty"
	}
	return 200, ""
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

func (c *NotifyType) SendMail(notify lib.NotifyContent) error {
	fmt.Println("sand mail:", notify)
	return mail.Send(notify)
}

func (c *NotifyType) SendSms(notify lib.NotifyContent) error {
	fmt.Println("sand sms:", notify)
	return sms.Send(notify)
}

func (c *NotifyType) SendHipchat(notify lib.NotifyContent) error {
	fmt.Println("sand hipchat:", notify)
	return hipchat.Send(notify)
}

func DoNotify(notify NotifyType) error {
	var notifyContent lib.NotifyContent
	notifyContent.Subject = notify.Subject
	notifyContent.Content = notify.Content
	notifyContent.Recieve = notify.Recieve
	notifyContent.PhoneList = notify.PhoneList  // ??? todo

	channelList := strings.Split(notify.Channel, ";")
	sort.Strings(channelList)
	notifyChannel := RemoveDuplicatesAndEmpty(channelList)
	for _, channel := range notifyChannel {
		switch channel {
		case "mail":
			if err := notify.SendMail(notifyContent); err != nil {
				return err
			}
		case "sms":
			if err := notify.SendSms(notifyContent); err != nil {
				return err
			}
		case "hipchat":
			if err := notify.SendHipchat(notifyContent); err != nil {
				return err
			}
		default:
			fmt.Println("unknown notify channel and skip:", channel)
		}

	}
	return nil
}
