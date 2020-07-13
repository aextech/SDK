package types

import (
	"net/url"
)

//认证
type Auth struct {
	*Cmd
	Action string `json:"action"`
	Key    string `json:"key"`
	Time   int64  `json:"time"`
	Md5    string `json:"md5"`
}

var U url.URL
