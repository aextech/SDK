package lib

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (h *Hup) read() {
	defer func() {
		_ = h.conn.Close()
		h.reStatus <- true
	}()
	h.conn.SetReadLimit(wsMaxMessageSize)
	_ = h.conn.SetWriteDeadline(time.Now().Add(wsPongWait))
	h.conn.SetPongHandler(func(string) error {
		_ = h.conn.SetReadDeadline(time.Now().Add(wsPongWait))
		return nil
	})
	for {
		_, msg, err := h.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			}
			log.Println("读取ws失败#3" + err.Error())
			break
		}
		h.resp <- msg
	}
}
