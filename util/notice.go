package util

import "novel-update-notice/config"

var (

	// 发送微信消息相关
	toTag             string
	tokenAPI          string
	messageAPI        string
	corpID            string
	warningAppAgentID int64
	warningAppSecret  string

	// 发送邮件相关
	emailType     string
	to            string
	selfHostFrom  string
	sendCloudFrom string
	fromName      string

	apiUser string
	apiKey  string
	apiUrl  string

	host     string
	username string
	port     int
	password string

	// 通知类型
	noticeType string
)

func getConfig() {
	iniConf := config.GetConfig()
	toTag = iniConf.String("weixin::toTag")
	tokenAPI = iniConf.String("weixin::accessTokenAPI")
	messageAPI = iniConf.String("weixin::sendMessageAPIURL")
	corpID = iniConf.String("weixin::corpID")
	warningAppAgentID, _ = iniConf.Int64("weixin::warningAppAgentID")
	warningAppSecret = iniConf.String("weixin::warningAppSecret")

	emailType = iniConf.String("email::type")
	to = iniConf.String("email::to")

	apiUser = iniConf.String("email_sendcloud::api_user")
	apiKey = iniConf.String("email_sendcloud::api_key")
	apiUrl = iniConf.String("email_sendcloud::api_url")
	fromName = iniConf.String("email_sendcloud::fromName")
	sendCloudFrom = iniConf.String("email_sendcloud::from")

	selfHostFrom = iniConf.String("email_selfhost::from")
	host = iniConf.String("email_selfhost::host")
	username = iniConf.String("email_selfhost::username")
	port, _ = iniConf.Int("email_selfhost::port")
	password = iniConf.String("email_selfhost::password")

	noticeType = iniConf.String("notice::type")
}

func init() {
	getConfig()
}

// Message 内容
type Message struct {
	SelfHostFrom  string
	SendCloudFrom string
	FromName      string
	To            []string
	Cc            []string
	Subject       string
	ContentType   string
	Content       string
	Attach        string
}

type Notice interface {
	SendMessage() error
}

func SendMessage(s Notice) error {
	return s.SendMessage()
}

func Send(subject, content string) error {
	var err error
	switch noticeType {
	case "weixin":
		err = SendMessage(NewSendWeixin(content))
	case "mail":
		err = SendMessage(NewSendMail(subject, content))
	}
	return err
}
