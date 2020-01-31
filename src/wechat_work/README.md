# Wechat Work robot client

## Build

```
git clone https://github.com/x-actions/wechat-work.git
cd wechat-work
export GOPATH=`pwd`
cd src/wechat_work/
GOOS=linux GOARCH=amd64 go build -tags netgo
```

## Usage

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

demo:

```
./wechat_work -webhookKey <webhookKey> -msgtype markdown -content "Build Done."
```

## Doc

- [https://work.weixin.qq.com/api/doc/90000/90136/91770](https://work.weixin.qq.com/api/doc/90000/90136/91770)
