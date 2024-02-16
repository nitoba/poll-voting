package ws

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type Client struct {
	//User ID
	Id string
	//Connected socket
	Socket *websocket.Conn
	//Message
	Send chan []byte
}

type ClientManager struct {
	//The client map stores and manages all long connection clients, online is TRUE, and those who are not there are FALSE
	Clients map[*Client]bool
	//Web side MESSAGE we use Broadcast to receive, and finally distribute it to all clients
	// Broadcast chan []byte
	//Newly created long connection client
	Register chan *Client
	//Newly canceled long connection client
	Unregister chan *Client
}

// Will formatting Message into JSON
type Message struct {
	//Message Struct
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var Manager = ClientManager{
	// Broadcast:  make(chan []byte),
	Register:   make(chan *Client),
	Unregister: make(chan *Client),
	Clients:    make(map[*Client]bool),
}

func (manager *ClientManager) Start() {
	for {
		select {
		//If there is a new connection access, pass the connection to conn through the channel
		case conn := <-manager.Register:
			//Set the client connection to true
			manager.Clients[conn] = true
			//Format the message of returning to the successful connection JSON
			jsonMessage, _ := json.Marshal(&Message{Content: "/A new socket has connected. "})
			//Call the client's send method and send messages
			manager.send(jsonMessage, nil)
			//If the connection is disconnected
		case conn := <-manager.Unregister:
			//Determine the state of the connection, if it is true, turn off Send and delete the value of connecting client
			if _, ok := manager.Clients[conn]; ok {
				close(conn.Send)
				delete(manager.Clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected. "})
				manager.send(jsonMessage, conn)
			}
			//broadcast
			// case message := <-manager.Broadcast:
			// 	//Traversing the client that has been connected, send the message to them
			// 	for conn := range manager.Clients {
			// 		select {
			// 		case conn.Send <- message:
			// 		default:
			// 			close(conn.Send)
			// 			delete(manager.Clients, conn)
			// 		}
			// 	}
		}
	}
}

// Define the send method of client management
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.Clients {
		//Send messages not to the shielded connection
		if conn != ignore {
			conn.Send <- message
		}
	}
}

func (c *Client) Write() {
	defer func() {
		_ = c.Socket.Close()
	}()

	for {
		select {
		//Read the message from send
		// Read the message from send
		case message, ok := <-c.Send:
			//If there is no message
			// If there is no message
			if !ok {
				// Close the websocket connection
				_ = c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//Write it if there is news and send it to the web side
			_ = c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func GetConnection() *websocket.Conn {
	return nil
}
