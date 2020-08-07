package lib

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

func (h *Hub) write() {
	ping := time.NewTicker(wsPingPeriod)
	defer func() {
		_ = h.conn.Close()
		ping.Stop()
	}()
	for {
		select {
		case msg, ok := <-h.req:
			msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
			if !ok {
				_ = h.conn.WriteMessage(websocket.CloseMessage, []byte{})
			}
			w, _ := h.conn.NextWriter(websocket.TextMessage)
			_, _ = w.Write(msg)

			if err := w.Close(); err != nil {
				log.Println("Err#1:" + err.Error())
				return
			}
		case <-ping.C:
			_ = h.conn.SetWriteDeadline(time.Now().Add(wsPingPeriod))
			if err := h.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Println("ws#2 断开:" + err.Error())
				return
			}
		}
	}
}
