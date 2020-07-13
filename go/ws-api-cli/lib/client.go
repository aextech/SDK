package lib

import (
	"flag"
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

var host = flag.String("H", "api.aex.zone", "input host")

var scheme = flag.String("s", "wss", "input scheme")

var path = flag.String("p", "/v3", "input path")

var key = flag.String("key", "", "input Key")

var skey = flag.String("skey", "", "input Skey")

var id = flag.String("id", "", "input userId")

func Client() {
	Parse()

	Hub := &Hub{
		Url:      types.U,
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

func Parse() {
	flag.Parse()

	if *host == "" {
		log.Fatal("host is not allowed to be empty")
	}

	if *scheme == "" {
		log.Fatal("scheme is not allowed to be empty")
	}

	if *path == "" {
		log.Fatal("path is not allowed to be empty")
	}

	types.U = url.URL{Scheme: *scheme, Host: *host, Path: *path}

	types.Key = *key
	types.Skey = *skey
	types.Id = *id
}
