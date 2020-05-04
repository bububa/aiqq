package aiqq

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Ret  int             `json:"ret,omitempty"`
	Msg  string          `json:"msg,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

func (this *Response) IsError() bool {
	return this.Ret != 0
}

func (this *Response) Error() string {
	return fmt.Sprintf("RET:%d, MSG:%s", this.Ret, this.Msg)
}
