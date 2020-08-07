package lib

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (h *Hub) read() {
	defer func() {
		_ = h.conn.Close()
		h.reStatus <- true
	}()
	h.conn.SetReadLimit(wsMaxMessageSize)
	_ = h.conn.SetWriteDeadline(time.Now().Add(wsPongWait))
	_ = h.conn.SetReadDeadline(time.Now().Add(wsPongWait))
	for {
		_, msg, err := h.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			log.Println("读取ws失败#3" + err.Error())
			break
		}
		_ = h.conn.SetReadDeadline(time.Now().Add(wsPongWait))
		_ = h.conn.SetWriteDeadline(time.Now().Add(wsPongWait))
		if string(msg) != "pong" {
			h.resp <- msg
		}
	}
}
