package types

import "time"

const (
	CmdConnSucc = 0 //连接成功
	CmdAuth     = 4 //签名认证
)

//重连频率
const ReConnLimit = 500 * time.Millisecond
