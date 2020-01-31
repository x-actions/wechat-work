package robot

const (
	msgTypeText     = "text"
	msgTypeMarkdown = "markdown"
)

type textMessage struct {
	MsgType string     `json:"msgtype"`
	Text    textParams `json:"text"`
}

type textParams struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty`
}

type markdownMessage struct {
	MsgType  string         `json:"msgtype"`
	Markdown markdownParams `json:"markdown"`
}

type markdownParams struct {
	Content  string `json:"content"`
}
