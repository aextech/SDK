package lib

import (
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
	"ws-api-cli/handle"
	"ws-api-cli/types"
)

type Hub struct {
	Url      url.URL
	req      chan []byte
	resp     chan []byte
	conn     *websocket.Conn
	reStatus chan bool
	reCount  int
}

func Client(host, path, scheme string) {

	u := url.URL{Scheme: scheme, Host: host, Path: path}

	Hub := &Hub{
		Url:      u,
		req:      make(chan []byte, 100),
		resp:     make(chan []byte, 100),
		conn:     &websocket.Conn{},
		reStatus: make(chan bool),
		reCount:  -1,
	}

	go Hub.Create(&handle.Body{
		Receive: make(chan []byte, 100),
		Send:    make(chan []byte, 100),
	})

	re := time.NewTicker(types.ReConnLimit)

	for {
		ok := Hub.connect()

		Hub.reCount++

		if !ok {
			log.Println("尝试重连...")
			<-re.C
			continue
		}

		go Hub.read()

		go Hub.write()

		<-Hub.reStatus
	}
}

func (h *Hub) Create(handle Handle) {
	b := handle.GetBody()

	go b.Work()

	for {
		select {
		case msg := <-b.Send:
			h.req <- msg

		case msg := <-h.resp:
			b.Receive <- msg
		}
	}
}
