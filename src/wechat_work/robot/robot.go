package robot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Roboter is the interface implemented by Robot that can send multiple types of messages.
type Roboter interface {
	SendText(content string, mentionedList, mentionedMobileList []string) error
	SendMarkdown(content string) error
}

// Robot represents a wechat wrok custom robot that can send messages to groups.
type Robot struct {
	webURL string
}

// NewRobot returns a roboter that can send messages.
func NewRobot(webURL string) Roboter {
	return &Robot{webURL: webURL}
}

// SendText send a text type message.
func (r Robot) SendText(content string, mentionedList, mentionedMobileList []string) error {
	return r.send(&textMessage{
		MsgType: msgTypeText,
		Text: textParams{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	})
}

// SendMarkdown send a markdown type message.
func (r Robot) SendMarkdown(content string) error {
	return r.send(&markdownMessage{
		MsgType: msgTypeMarkdown,
		Markdown: markdownParams{
			Content: content,
		},
	})
}

type wechatWrokResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (r Robot) send(msg interface{}) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	webURL := r.webURL
	resp, err := http.Post(webURL, "pplication/json", bytes.NewReader(m))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var res wechatWrokResponse
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}
	if res.Errcode != 0 {
		return fmt.Errorf("wechatWrokrobot send failed: %v", res.Errmsg)
	}

	fmt.Println("Wechat Work API Response is", res)

	return nil
}
