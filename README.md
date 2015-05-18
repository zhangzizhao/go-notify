##go-notify
- code in golang
- notify service:   notify msg through mail/sms/hipchat   (hipchat not done)
- mail: qq corp mailserver
- sms: self thrift protocol
- hipchat: todo
- web 框架： gin

###my habit
- build project : make build
- fmt:  make fmt
- init: make dep


###param
can notify by post reuqest, example for body param:

{
  "channel": "mail;sms", 
  "subject": "notify 发送-标题",
  "content": "hello <br>gin 发送- content",
  "receive": "zhangzizhao;zhangzizhao02",
  "phonelist": [13521022145]
}
- channel: notify channel, include mail/sms   and hipchat(todo)
- subject: mail subject
- content: mail/sms/hipchat msg
- recieve: mail/hipchat recieve
- phonelist: phone list for sms


###todo:
1. a server to turn name to phone, to don't neet to give phonelist
2. clear code, 优雅
