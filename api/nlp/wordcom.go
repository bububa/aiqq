package nlp

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/aiqq"
)

type WordcomRequest struct {
	Text string `json:"text"` // 待分析文本
}

func (this *WordcomRequest) Path() string {
	return "nlp/nlp_wordcom"
}

func (this *WordcomRequest) Method() string {
	return "post"
}

func (this *WordcomRequest) Values() url.Values {
	values := url.Values{}
	values.Set("text", this.Text)
	return values
}

type WordcomResponse struct {
	Text      string     `json:"text,omitempty"`       // API请求中的待分析文本
	Intent    IntentType `json:"intent,omitempty"`     // 意图编码
	ComTokens []ComToken `json:"com_tokens,omitempty"` // 成分列表
}

type ComToken struct {
	Type ComType `json:"com_type,omitempty"` // 成分编码
	Word string  `json:"com_word,omitempty"` // 成分分词
}

func Wordcom(clt *aiqq.Client, text string) (*WordcomResponse, error) {
	resp, err := clt.Do(&WordcomRequest{Text: text})
	if err != nil {
		return nil, err
	}
	var ret WordcomResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
