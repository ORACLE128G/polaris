package connection

import (
	"errors"
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	Conn      *websocket.Conn
	In        chan []byte
	Out       chan []byte
	CloseChan chan byte
	lock      sync.Mutex
	Closed    bool
}

func InitConnetion(ws *websocket.Conn) (conn *Connection, err error) {
	in := make(chan []byte, 1000)
	c := &Connection{
		Conn:      ws,
		In:        in,
		Out:       make(chan []byte, 1000),
		CloseChan: make(chan byte, 1),
	}
	go c.ReadLoop()
	go c.WriteLoop()
	return c, nil
}

func (c *Connection) ReadMsg() (in []byte, err error) {
	select {
	case in = <-c.In:
	case <-c.CloseChan:
		return nil, errors.New("connection was closed")
	}
	return in, nil
}

func (c *Connection) WriteMsg(out []byte) error {
	select {
	case c.Out <- out:
	case <-c.CloseChan:
		return errors.New("connection was closed")
	}
	return nil
}

func (c *Connection) Close() error {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.Closed {
		close(c.CloseChan)
		c.Closed = true
	}
	return c.Conn.Close()
}

func (c *Connection) ReadLoop() {
	for {
		if _, data, err := c.Conn.ReadMessage(); err != nil {
			goto ERR
		} else {
			select {
			case c.In <- data:
			case <-c.CloseChan:
				goto ERR
			}
		}
	}
ERR:
	c.Close()
}

func (c *Connection) WriteLoop() {
	for {
		select {
		case data := <-c.Out:
			if err := c.Conn.WriteMessage(websocket.TextMessage, data); err != nil {
				break
			}
		case <-c.CloseChan:
			break

		}

	}
	c.Close()
}
