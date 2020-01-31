# Github Action for Sending Wechat Wrok robot messages.

![](https://github.com/x-actions/wechat/workflows/wechat/badge.svg)

Github Action for Sending Wechat Wrok robot messages.

## Environment Variables

- WECHAT_WEBHOOK_KEY: wechat webhook key

## How to Use

```
    - name: Sending Wechat Work Message
      uses: x-actions/wechat-work@master
      env:
        WECHAT_WEBHOOK_KEY: ${{ secrets.WECHAT_WEBHOOK_KEY }}
        MSGTYPE: markdown
        CONTENT: |
          # Wechat Work Noti
          > ^_^
```

## Options

```
./wechat_work
Usage of ./wechat_work:
  -content string
    	content of message (default "This is a text")
  -mentioned string
    	@user, like:"some","@all"
  -msgtype string
    	message type (text, markdown, image, news) (default "markdown")
  -webhookKey string
    	robot webhook access token
```

## Doc

- [https://work.weixin.qq.com/api/doc/90000/90136/91770](https://work.weixin.qq.com/api/doc/90000/90136/91770)
