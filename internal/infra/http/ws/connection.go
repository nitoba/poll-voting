package ws

import "github.com/gorilla/websocket"

var conn *websocket.Conn

func GetConnection() *websocket.Conn {
	return conn
}

func SetConnection(c *websocket.Conn) {
	conn = c
}
