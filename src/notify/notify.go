package main

import (
        "test/gin/lib"
        "fmt"
        "test/gin/mail"
)

type NotifyType struct {
	Mail    bool   `json:"mail"`
	Sms     bool   `json:"sms"`
	Subject string `json:"subject"`
	Content string `json:"content"`
	Recieve string `json:"receive"`
}

type Mail struct {
    Subject string  `json:"subject"`
    Content string  `json:"content"`
    Recieve string  `json:"recieve"`
}

type Sms struct {
    Subject string  `json:"subject"`
    Content string  `json:"content"`
    Recieve string  `json:"recieve"`
}

func (c *NotifyType)SendMail(notify lib.NotifyContent) {
    fmt.Println("sand mail:", notify)
    mail.Send(notify)
}

func (c *NotifyType)SendSms(notify lib.NotifyContent) {
    fmt.Println("sand sms:", notify)
}

func DoNotify(notify NotifyType) error {
    var notifyContent lib.NotifyContent
    notifyContent.Subject = notify.Subject
    notifyContent.Content = notify.Content
    notifyContent.Recieve = notify.Recieve
    if notify.Mail == true {
        notify.SendMail(notifyContent)
    }
    if notify.Sms == true {
        notify.SendSms(notifyContent)
    }
    return nil
}
