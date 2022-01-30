package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	send   chan []byte
	routes WSRouteMap
}

type Message struct {
	Type string `json:"type"`
}

func (c *Client) writer() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

func (c *Client) reader() {
	c.conn.SetReadLimit(512)

	defer func() {
		c.conn.Close()
	}()

	for {
		var parsed Message
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) || websocket.IsCloseError(err, websocket.CloseAbnormalClosure) {
				log.Println("user disconnected")
			}
			break
		}
		err = json.Unmarshal(message, &parsed)
		if err != nil {
			log.Println(err)
			break
		}

		handler, ok := c.routes[parsed.Type]
		if ok {
			handler(message, c.conn.WriteJSON)
		} else {
			fmt.Println("Uknown message")
		}
	}
}
