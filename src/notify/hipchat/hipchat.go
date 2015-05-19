package hipchat

import (
	"fmt"
	"net/http"
	"strings"

	"notify/lib"
)

func Send(notify lib.NotifyContent) error {
	usrList := strings.Split(notify.Recieve, ";")
	for _, oneUser := range usrList {
		if err := sendMsgToOneUser(notify.Content, oneUser); err != nil {
			return err
		}
	}
    return nil
}

func sendMsgToOneUser(msg string, usr string) error {
	urlSendMsg := fmt.Sprintf(lib.HipchatUrl, usr, lib.HipchatToken)
	res, err := http.Post(urlSendMsg, "text/plain", strings.NewReader(msg))
	if err != nil {
		return err
	}
	defer res.Body.Close()
    return nil
}
