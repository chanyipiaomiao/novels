package util

import (
	"fmt"
	"github.com/chanyipiaomiao/hltool"
	"github.com/levigross/grequests"
	"github.com/tidwall/gjson"
	"strings"
)

type SelfHostMail struct {
	Username string
	Password string
	Port     int
	Host     string
}

type SendCloudMail struct {
	APIUser  string
	APIKey   string
	APIUrl   string
	Template string
}

// EmailClient 发送客户端
type EmailClient struct {
	Type    string
	Message *Message
	SelfHostMail
	SendCloudMail
}

func newEmailClient(subject, content string) *EmailClient {
	message := &Message{
		Subject:       subject,
		Content:       content,
		ContentType:   "text/html",
		To:            []string{to},
		SelfHostFrom:  selfHostFrom,
		SendCloudFrom: sendCloudFrom,
		FromName:      fromName,
	}
	client := &EmailClient{Message: message, Type: emailType}

	if emailType == "sendcloud" {
		client.APIUser = apiUser
		client.APIKey = apiKey
		client.APIUrl = apiUrl
	} else {
		client.Host = host
		client.Username = username
		client.Port = port
		client.Password = password
	}
	return client
}

type SendMail struct {
	Client *EmailClient
}

func NewSendMail(subject, content string) *SendMail {
	return &SendMail{
		Client: newEmailClient(subject, content),
	}
}

func (s *SendMail) SendMessage() error {
	var err error
	switch s.Client.Type {
	case "sendcloud":
		err = bySendCloudMail(s.Client)
	case "selfhost":
		err = bySelfHostMail(s.Client)
	default:
		err = fmt.Errorf("no support %s", s.Client.Type)
	}
	return err
}

func bySendCloudMail(email *EmailClient) error {
	o := &grequests.RequestOptions{
		Data: map[string]string{
			"apiUser":  email.APIUser,
			"apiKey":   email.APIKey,
			"from":     email.Message.SendCloudFrom,
			"fromName": email.Message.FromName,
			"to":       strings.Join(email.Message.To, ";"),
			"subject":  email.Message.Subject,
			"html":     email.Message.Content,
		},
	}
	resp, err := grequests.Post(email.APIUrl, o)
	if err != nil {
		return err
	}

	respS := resp.String()
	r := gjson.Get(respS, "statusCode")
	if r.String() != "200" {
		return fmt.Errorf("call SendCloud error: %s", gjson.Get(respS, "message").String())
	}
	return nil
}

func bySelfHostMail(email *EmailClient) error {
	message := hltool.NewEmailMessage(email.Message.SelfHostFrom, email.Message.Subject, email.Message.ContentType,
		email.Message.Content, "", email.Message.To, []string{})
	client := hltool.NewEmailClient(email.Host, email.Username, email.Password, email.Port, message)
	_, err := hltool.SendMessage(client)
	if err != nil {
		return err
	}
	return nil
}
