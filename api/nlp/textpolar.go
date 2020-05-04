package nlp

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/aiqq"
)

type TextpolarRequest struct {
	Text string `json:"text"` // 待分析文本
}

func (this *TextpolarRequest) Path() string {
	return "nlp/nlp_textpolar"
}

func (this *TextpolarRequest) Method() string {
	return "post"
}

func (this *TextpolarRequest) Values() url.Values {
	values := url.Values{}
	values.Set("text", this.Text)
	return values
}

type TextpolarResponse struct {
	Text  string    `json:"text,omitempty"`  // API请求中的待分析文本
	Polar PolarType `json:"polar,omitempty"` // 情感编码
	Confd float64   `json:"confd,omitempty"` // 置信度
}

func Textpolar(clt *aiqq.Client, text string) (*TextpolarResponse, error) {
	resp, err := clt.Do(&TextpolarRequest{Text: text})
	if err != nil {
		return nil, err
	}
	var ret TextpolarResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
