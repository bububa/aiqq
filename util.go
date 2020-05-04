package aiqq

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/nu7hatch/gouuid"
	"io"
	"net/url"
	"sort"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}

func Nonce() string {
	token, _ := uuid.NewV4()
	return Md5(token.String())
}

func SortParamters(params url.Values) []string {
	var keys []string
	for k, _ := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
