package main

import (
	// built-in packages
	"fmt"
	"log"
	"strconv"
	"strings"

	// local packages
	"go-catchmind/src/models"
	"go-catchmind/src/routes"
	"go-catchmind/src/utils"

	// external packages
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

const DEFAULT_PORT int = 3000;

func main() {
	// Configure
	port := utils.GetPort()
	
	// Web Socket
	server := socketio.NewServer(nil)
	defer server.Close()

	// Regist Routes
	// canvas
	canvasState := models.CanvasState{[][]int{}}
	routes.RegistCanvasRoutes(server, &canvasState)

	// Run Websocket
	go func() {
		if err := server.Serve(); err != nil {
			fmt.Println("HTTP ERROR: ", err)
			log.Fatal("Web Socket ERROR: ", err)
		}
	}()

	// HTTP
	router := gin.New()
	router.GET("/socket.io/*any", gin.WrapH(server))
	router.POST("/socket.io/*any", gin.WrapH(server))
	router.Use(static.Serve("/", static.LocalFile("./client/build", true))) 

	// run server
	portStrs := []string{":", strconv.Itoa(port)}
	portStr := strings.Join(portStrs, "")
	if err := router.Run(portStr); err != nil {
		// failed run
		fmt.Println("HTTP ERROR: ", err)
		log.Fatal("HTTP ERROR: ", err)
	} else {
		// success run
		fmt.Println("run socket.io go server port:", port);
	}
}

