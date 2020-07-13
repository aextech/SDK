package types

//认证
type Auth struct {
	*Cmd
	Action string `json:"action"`
	Key    string `json:"key"`
	Time   int64  `json:"time"`
	Md5    string `json:"md5"`
}

//最新公共成交数据
type PublicTrade struct {
	*Cmd
	Action string `json:"action"`
	Symbol string `json:"symbol"`
}
