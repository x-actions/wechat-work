package main

import (
	"flag"
	"fmt"
	"os"

	"wechat_work/robot"
)

var msgtype string
var webhookKey string
var content string

var mentioned string

func init() {
	flag.StringVar(&msgtype, "msgtype", "markdown", "message type (text, markdown, image, news)")
	flag.StringVar(&webhookKey, "webhookKey", "", "robot webhook access token")
	flag.StringVar(&content, "content", "This is a text", "content of message")
	flag.StringVar(&mentioned, "mentioned", "", "@user, like:\"some\",\"@all\"")
	flag.Parse()
}

func usage() {
	flag.Usage()
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	var err error

	webURL := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + webhookKey
	r := robot.NewRobot(webURL)

	switch msgtype {
	case "text":
		err = r.SendText(content, []string{}, []string{})
	case "markdown":
		err = r.SendMarkdown(content)
	case "image":
	case "news":
	}
	if err != nil {
		fmt.Println("err:", err)
	}
}
