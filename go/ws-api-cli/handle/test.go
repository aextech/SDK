package handle

import (
	"encoding/json"
	"github.com/wangxudong123/assist"
	"log"
	"strconv"
	_func "ws-api-cli/func"
	"ws-api-cli/types"
)

//---------------------------------
//业务逻辑请在work()方法里面写
//1.msg<-b.Receive接收msg
//2.b.send<-msg发送msg
//如此简单,敬请享用吧,祝您早日暴富 :)
//---------------------------------
type Body struct {
	Receive chan []byte
	Send    chan []byte
	//Test map[string]string   //自定义属性在Work里面初始化
}

//TODO 这里不要改动
func (b *Body) GetBody() *Body {
	return b
}

//工作方法
func (b *Body) Work() {
	//b.Test = make(map[string]string)

	resp := types.RespCmd{}
	for {
		select {
		case msg := <-b.Receive: //接收消息
			log.Println("Receive: " + string(msg))

			_ = json.Unmarshal(msg, &resp)

			//处理消息
			b.option(resp)
		}
	}
}

//接收消息的操作
func (b *Body) option(resp types.RespCmd) {
	switch resp.Cmd.Cmd {
	case types.CmdConnSucc: //连接成功
		//-------------
		//连接成功的操作
		//-------------
		//认证
		//b.Auth()
		//b.PublicTrade("btc_cnc")
		b.PublicTrade("btc_cnc")
		b.PublicOrder("btc_cnc")
		b.PublicKline("btc_cnc@1min")
		b.PublicMarket("all")
		break
	case types.CmdAuth: //登录认证
		if codeStatus(resp.Code) == false {
			//认证失败
			log.Println("auth fail [code]:" + strconv.Itoa(resp.Code))
			break
		}

		//-------------
		//认证成功的操作
		//-------------

		//订阅最新成交数据

		break
	case types.CmdPublicTrade:
		if codeStatus(resp.Code) == false {
			//订阅最新成交数据失败
			log.Println("public-trade fail [code]:" + strconv.Itoa(resp.Code))
			break
		}

		//case ...
	default:
		//log.Println("不存在的命令字")
		break
	}
}

//认证
func (b *Body) Auth() {
	auth := types.Auth{
		Cmd:    &types.Cmd{Cmd: types.CmdAuth},
		Action: "login",
	}

	var time string
	auth.Md5, auth.Key, time = _func.Sign()
	auth.Time = assist.StringToInt64(time)

	b.send(auth)
}

//请求参数symbol都支持多个订阅，用逗号隔开连接
func (b *Body) PublicTrade(symbol string) {
	b.send(types.Types{
		Cmd:    &types.Cmd{Cmd: types.CmdPublicTrade},
		Action: "sub",
		Symbol: symbol,
	})
}

//请求参数symbol都支持多个订阅，用逗号隔开连接
func (b *Body) PublicOrder(symbol string) {
	b.send(types.Types{
		Cmd:    &types.Cmd{Cmd: types.CmdPublicOrder},
		Action: "sub",
		Symbol: symbol,
	})
}

//请求参数symbol都支持多个订阅，用逗号隔开连接
func (b *Body) PublicKline(symbol string) {
	b.send(types.Types{
		Cmd:    &types.Cmd{Cmd: types.CmdPublicKline},
		Action: "sub",
		Symbol: symbol,
	})
}

//请求参数symbol都支持多个订阅，用逗号隔开连接
func (b *Body) PublicMarket(symbol string) {
	b.send(types.Types{
		Cmd:    &types.Cmd{Cmd: types.CmdPublicMarket},
		Action: "sub",
		Symbol: symbol,
	})
}

//return bool
//true is ok
//false is fail
func codeStatus(code int) bool {
	s := strconv.Itoa(code)
	if string(s[0]) == "0" {
		return false
	}

	if assist.StringToInt(string(s[0]))%2 == 0 {
		return true
	}
	return false
}

func (b *Body) send(v interface{}) {
	msg, _ := json.Marshal(&v)
	log.Println("Send: " + string(msg))
	b.Send <- msg
}
