package types

import "time"

const (
	CmdConnSucc = 0  //连接成功
	CmdAuth     = 99 //签名认证
)

//重连频率
const ReConnLimit = 500 * time.Millisecond
