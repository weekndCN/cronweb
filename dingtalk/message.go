package dingtalk

import (
	"fmt"
	"strings"
)

// https://ding-doc.dingtalk.com/doc#/serverapi2/qf2nxq
// Msg for shorten message
const (
	msgText       = "text"       // text type
	msgLink       = "link"       // link type
	msgMarkDown   = "markdown"   // markdown type
	msgActionCard = "actionCard" // ActionCard type
)

type atReceiver struct {
	AtMobiles []string `json:"atMobiles,omitemty"`
	IsAtAll   bool     `json:"isAtAll,omitemty"`
}

type textMsg struct {
	MsgType string      `json:"msgtype"`
	Text    textContent `json:"text"`
	At      atReceiver  `json:"at"`
}

type textContent struct {
	Content string `json:"content"`
}

type linkMsg struct {
	MsgType string      `json:"msgtype"`
	Link    linkContent `json:"link"`
}

type linkContent struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	MsgURL string `json:"messageUrl"`
	PicURL string `json:"picUrl,omitempty"`
}

type markdownMsg struct {
	MsgType  string          `json:"msgtype"`
	Markdown markdownContent `json:"markdown"`
	At       atReceiver      `json:"at"`
}

type markdownContent struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type actionCardMsg struct {
	MsgType    string            `json:"msgtype"`
	ActionCard actionCardContent `json:"actionCard"`
}

type actionCardContent struct {
	Title          string `json:"title"`
	Text           string `json:"text"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
	BtnOrientation string `json:"btnOrientation,omitempty"`
	HideAvatar     string `json:"hideAvatar,omitempty"`
}

// MsgText dingtalk msg
func MsgText(color, title, url, result, start, log string, duration int64, statuscode int) string {
	/*
		red: #FF0000
		green: #0DAD51
		yellow: #FF8000
	*/
	var text []string
	// title
	text = append(text, fmt.Sprintf("# <font color=%s>%s</font>", color, title))
	// url
	text = append(text, fmt.Sprintf("---"))
	text = append(text, fmt.Sprintf("> **URL: %s**", url))
	// result
	text = append(text, fmt.Sprintf("> **结果:** **<font color=%s> %s</font>**", color, result))
	text = append(text, fmt.Sprintf("> **状态码:** **<font color=%s> %d</font>**", color, statuscode))

	// start
	text = append(text, fmt.Sprintf("> **开始时间: %s**", start))
	// duration
	text = append(text, fmt.Sprintf("> **耗时: %d millionseconds**", duration))
	text = append(text, fmt.Sprintf("---"))
	// log
	text = append(text, fmt.Sprintf("**返回日志:** \n%s", strings.Replace(log, "#", "-", -1)))
	dingmsg := strings.Join(text, "\n\n")
	return fmt.Sprintf("%s", dingmsg)
}
