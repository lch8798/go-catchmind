package routes

import (
	"go-catchmind/src/controllers"
	"go-catchmind/src/models"

	socketio "github.com/googollee/go-socket.io"
)

func RegistCanvasRoutes(server *socketio.Server, state *models.CanvasState) {
	server.OnConnect("/", controllers.OnConnect)
	server.OnEvent("/", "drawInit", controllers.DrawInit(state))
	server.OnEvent("/", "paintInit", controllers.PaintInit(state))
	server.OnEvent("/", "draw", controllers.Draw(state));
	server.OnError("/", controllers.OnError)
	server.OnDisconnect("/", controllers.OnDisconnect)
}