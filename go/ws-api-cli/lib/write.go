package lib

import (
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
			_ = h.conn.SetWriteDeadline(time.Now().Add(wsWriteWait))
			_ = h.conn.SetReadDeadline(time.Now().Add(wsPongWait))
			//msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
			if !ok {
				_ = h.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			w, err := h.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println("Err#2:" + err.Error())
				return
			}
			_, err = w.Write(msg)
			if err != nil {
				log.Println("Err#3:" + err.Error())
				return
			}

			if err := w.Close(); err != nil {
				log.Println("Err#1:" + err.Error())
				return
			}
		case <-ping.C:
			if !h.ping() {
				return
			}
			_ = h.conn.SetWriteDeadline(time.Now().Add(wsWriteWait))
			_ = h.conn.SetReadDeadline(time.Now().Add(wsPongWait))
		}
	}
}

func (h *Hub) ping() bool {
	w, err := h.conn.NextWriter(websocket.TextMessage)
	if err != nil {
		log.Println("pingErr#3:" + err.Error())
		return false
	}
	_, err = w.Write([]byte("ping"))
	if err != nil {
		log.Println("pingErr#3:" + err.Error())
		return false
	}

	if err := w.Close(); err != nil {
		log.Println("pingErr#1:" + err.Error())
		return false
	}
	return true
}
