package main

import (
	// built-in packages
	"fmt"
	"log"
	"strconv"
	"strings"

	// local packages
	"go-catchmind/src/utils"

	// external packages
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
)

const DEFAULT_PORT int = 8000;

func main() {
	// Configure
	port := utils.GetPort()
	

	// Web Socket
	server := socketio.NewServer(nil)
	defer server.Close()

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println(s)
		return nil
	})

	server.OnEvent("/", "drawInit", func(s socketio.Conn) {
		s.Emit("drawInit", "test")
		fmt.Println("drawInit")
	})

	// 실시간 랜더링
	server.OnEvent("/", "draw", func(s socketio.Conn, msg string) {
		// paint.push(data)
		// s.Emit("draw", data)
		fmt.Println(msg)
		fmt.Println("11")
	});

	// 캔버스 초기화
	// server.OnEvent('paintInit', (data) => { 
	// 	paint = []; 
	// 	io.emit('paintInit', paint);
	// });

	server.OnError("/", func(s socketio.Conn, err error) {
		log.Println("error:", err)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		log.Println("disconnect", reason)
	})

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

