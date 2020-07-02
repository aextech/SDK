package handle

import (
	"encoding/json"
	"log"
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
		case msg := <-b.Receive:
			log.Println("Receive: " + string(msg))

			_ = json.Unmarshal(msg, &resp)

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

		break
	case types.CmdAuth: //登录认证
		if resp.Code != 0 {
			//认证失败

			break
		}
		//-------------
		//认证成功的操作
		//-------------

		break

		//case ...
	default:
		log.Println("不存在的命令字")
	}
}

func (b *Body) send(v interface{}) {
	msg, _ := json.Marshal(&v)

	log.Println("Send: " + string(msg))
	b.Send <- msg
}
