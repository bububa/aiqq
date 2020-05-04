package aiqq

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bububa/aiqq/internal/debug"
)

type Client struct {
	http.Client
	appId  string
	appKey string
}

func NewClient(appId string, appKey string) *Client {
	return &Client{
		appId:  appId,
		appKey: appKey,
	}
}

func (this *Client) Do(r Request) (json.RawMessage, error) {
	values := this.RequestValues(r)
	var (
		resp *http.Response
		err  error
	)
	if r.Method() == "post" {
		requestUri := fmt.Sprintf("%s/%s", GATEWAY, r.Path())
		debug.DebugPrintPostMapRequest(requestUri, values)
		resp, err = this.PostForm(requestUri, values)
	} else {
		requestUri := fmt.Sprintf("%s/%s?%s", GATEWAY, r.Path(), values.Encode())
		debug.DebugPrintPostMapRequest(requestUri, values)
		resp, err = this.Get(requestUri)
	}

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ret Response
	err = debug.DecodeJSONHttpResponse(resp.Body, &ret)
	if err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, &ret
	}
	return ret.Data, nil
}

func (this *Client) RequestValues(r Request) url.Values {
	values := r.Values()
	ts := time.Now().Unix()
	values.Set("time_stamp", strconv.FormatInt(ts, 10))
	values.Set("nonce_str", Nonce())
	values.Set("app_id", this.appId)
	values.Set("sign", this.Sign(values))
	return values
}

func (this *Client) Sign(values url.Values) string {
	params := SortParamters(values)
	var vals []string
	for _, k := range params {
		v := values.Get(k)
		if v == "" {
			continue
		}
		vals = append(vals, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	vals = append(vals, fmt.Sprintf("app_key=%s", this.appKey))
	rawSign := strings.Join(vals, "&")
	return strings.ToUpper(Md5(rawSign))
}
