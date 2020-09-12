package types

import "time"

const (
	CmdConnSucc     = 0  //连接成功
	CmdAuth         = 99 //签名认证
	CmdPublicTrade  = 1  //最新公共成交数据
	CmdPublicOrder  = 3  //最新深度
	CmdPublicKline  = 2  //K线
	CmdPublicMarket = 4  //K线
)

//重连频率
const ReConnLimit = 500 * time.Millisecond
