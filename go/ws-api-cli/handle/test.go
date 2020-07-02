package handle

import "fmt"

type Body struct {
	Receive chan []byte
	Send    chan []byte
}

func (b *Body) GetBody() *Body {
	return b
}

func (b *Body) Work() {
	for {
		select {
		case msg := <-b.Receive:
			fmt.Println(string(msg))
		}
	}
}
