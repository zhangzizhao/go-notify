package mail

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"net/smtp"

	"notify/lib"
)

// reieve：  只填写统一账户用户名，;分割，自动加@后缀
func Send(notify lib.NotifyContent) {
	toEmail := fmt.Sprintf("%s@diditaxi.com.cn", notify.Recieve)
	subject := notify.Subject
	content := notify.Content

	header := make(map[string]string)
	header["From"] = "from odin" + "<" + lib.Email + ">"
	header["To"] = toEmail
	header["Subject"] = subject
	header["Content-Type"] = "text/html; charset=UTF-8"

	body := content

	message := ""
	for k, v := range header {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	auth := smtp.PlainAuth(
		"",
		lib.Email,
		lib.Password,
		lib.Host,
	)

	err := SendMailUsingTLS(
		fmt.Sprintf("%s:%d", lib.Host, lib.Port),
		auth,
		lib.Email,
		[]string{toEmail},
		[]byte(message),
	)
	if err != nil {
		panic(err)
	}
}

//return a smtp client
func Dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		log.Println("Dialing Error:", err)
		return nil, err
	}
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

func SendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	c, err := Dial(addr)
	if err != nil {
		log.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				log.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
