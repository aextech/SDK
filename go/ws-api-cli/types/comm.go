package types

import "net/url"

var Key string

var Skey string

var Id string

var U url.URL

type Cmd struct {
	Cmd int `json:"cmd"`
}
