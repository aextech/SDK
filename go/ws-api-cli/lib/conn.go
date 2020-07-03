package lib

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// 允许写入消息的时间。
	wsWriteWait = 10 * time.Second

	// 允许读取下一个消息的时间。
	wsPongWait = 60 * time.Second

	wsPingPeriod = (wsPongWait * 9) / 10

	//允许的最大消息。
	wsMaxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func (h *Hub) connect() bool {

	//开启压缩
	websocket.DefaultDialer.EnableCompression = true

	websocket.DefaultDialer.ReadBufferSize = 1024

	websocket.DefaultDialer.WriteBufferSize = 1024

	c, _, err := websocket.DefaultDialer.Dial(h.Url.String(), nil)
	if err != nil {
		log.Println("连接失败:" + err.Error())
		return false
	}

	log.Println("连接成功")

	h.conn = c

	return true
}
