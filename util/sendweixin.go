package util

import (
	"github.com/chanyipiaomiao/weixin-kit"
)

type SendWeixin struct {
	Content string
}

func NewSendWeixin(content string) *SendWeixin {
	return &SendWeixin{Content: content}
}

func (s *SendWeixin) SendMessage() error {
	message := &weixin.Message{
		MsgType: weixin.TEXT,       // 目前只支持发送文本消息
		ToTag:   toTag,             // ToTag 是在企业微信后台定义的标签ID，标签里面可以包含很多人 还有ToUser,ToParty参数 指定用户和部门ID
		AgentID: warningAppAgentID, // 企业应用的id，整型。可在应用的设置页面查看
		Safe:    0,                 // 表示是否是保密消息，0表示否，1表示是，默认0
		Text: &weixin.Text{
			Content: s.Content,
		},
	}

	client := &weixin.Client{
		AccessTokenAPI: tokenAPI,
		APIURL:         messageAPI,
		CorpID:         corpID,
		CorpSecret:     warningAppSecret,
		Message:        message,
	}

	_, err := client.SendMessage()
	if err != nil {
		return err
	}
	return nil
}
