package nlp

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/aiqq"
)

type TextchatRequest struct {
	Session  string `json:"session"`
	Question string `json:"question"` // 用户输入的聊天内容
}

func (this *TextchatRequest) Path() string {
	return "nlp/nlp_textchat"
}

func (this *TextchatRequest) Method() string {
	return "post"
}

func (this *TextchatRequest) Values() url.Values {
	values := url.Values{}
	values.Set("question", this.Question)
	session := this.Session
	if session == "" {
		session = aiqq.Nonce()
	}
	values.Set("session", session)
	return values
}

type TextchatResponse struct {
	Session string `json:"session,omitempty"`
	Answer  string `json:"answer,omitempty"`
}

func Textchat(clt *aiqq.Client, question string, session string) (*TextchatResponse, error) {
	resp, err := clt.Do(&TextchatRequest{Question: question, Session: session})
	if err != nil {
		return nil, err
	}
	var ret TextchatResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
