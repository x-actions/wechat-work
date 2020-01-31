package robot

import (
	"log"
)

func ExampleSendText() {
	// You should replace the webhook here with your own.
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx"
	robot := NewRobot(webhook)

	content := "我就是我,  @1352359XXXX 是不一样的烟火"
	atMobiles := []string{"1352359XXXX"}
	isAtAll := false

	err := robot.SendText(content, atMobiles, isAtAll)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleSendMarkdown() {
	// You should replace the webhook here with your own.
	webhook := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx"
	robot := NewRobot(webhook)

	title := "杭州天气"
	text := "#### 杭州天气  \n > 9度，@1352359XXXX 西北风1级，空气良89，相对温度73%\n\n > ![screenshot](http://i01.lw.aliimg.com/media/lALPBbCc1ZhJGIvNAkzNBLA_1200_588.png)\n  > ###### 10点20分发布 [天气](http://www.thinkpage.cn/) "
	atMobiles := []string{"1352359XXXX"}
	isAtAll := false

	err := robot.SendMarkdown(title, text, atMobiles, isAtAll)
	if err != nil {
		log.Fatal(err)
	}
}
