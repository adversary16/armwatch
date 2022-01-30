package ws

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WSRouteController func([]byte, func(interface{}) error)
type WSRouteMap map[string]WSRouteController

const (
	pingPeriod = 1 * time.Second
	writeWait  = 2 * time.Second
)

func CheckOrigin(r *http.Request) bool {
	return true
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     CheckOrigin,
}

func Controller(routeMap WSRouteMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
		client := &Client{conn: conn, send: make(chan []byte, 256), routes: routeMap}

		go client.writer()
		go client.reader()
	}
}
