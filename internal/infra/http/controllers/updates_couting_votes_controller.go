package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/nitoba/poll-voting/internal/infra/http/ws"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024 * 1024 * 1024,
	WriteBufferSize: 1024 * 1024 * 1024,
	//Solving cross-domain problems
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type UpdateCountingVotesController struct{}

func (ct *UpdateCountingVotesController) Handle(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	//Every connection will open a new client, client.id generates through UUID to ensure that each time it is different
	client := &ws.Client{Id: uuid.New().String(), Socket: conn, Send: make(chan []byte)}
	//Register a new link
	ws.Manager.Register <- client

	//Start the message to collect the news from the web side
	// go client.read()
	//Start the corporation to return the message to the web side
	go client.Write()
}

func NewUpdateCountingVotesController() *UpdateCountingVotesController {
	return &UpdateCountingVotesController{}
}
