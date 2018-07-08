package app

import (
	"github.com/kataras/iris/websocket"
	"fmt"
)

func (a *Application) SetupWebsockets(endpoint string, onConnection websocket.ConnectionFunc) {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(onConnection)

	a.Server.Get(endpoint, ws.Handler())
	a.Server.Any("/iris-ws.js", websocket.ClientHandler())
}

func handleConnection(c websocket.Connection) {
	c.On("watch", func(msg string) {
		c.Join(msg)
		// Print the message to the console, c.Context() is the iris's http context.
		fmt.Printf("%s sent: %s\n", c.Context().RemoteAddr(), msg)
		// Write message back to the client message owner:
		// c.Emit("chat", msg)
		c.To(msg).Emit("watch", "TEST")
	})

	c.OnLeave(func(roomName string) {
		fmt.Printf("\nConnection with ID: %s has been disconnected!", c.ID())
	})
}
