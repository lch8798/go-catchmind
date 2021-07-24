package controllers

import (
	"encoding/json"
	"fmt"
	"log"

	"go-catchmind/src/models"

	socketio "github.com/googollee/go-socket.io"
)

func OnConnect(s socketio.Conn) error {
	fmt.Println(s)
	return nil
}

// 초기 데이터
func DrawInit(state *models.CanvasState) func(socketio.Conn) {
	return func(s socketio.Conn) {
		s.Emit("paintInit", state.Canvas);
	}
}

// 캔버스 초기화
func PaintInit(state *models.CanvasState) func(socketio.Conn)  {
	return func(s socketio.Conn) {
		state.Canvas = [][]int{}
		s.Emit("paintInit", state.Canvas);
	}
}

// 실시간 랜더링
func Draw(state *models.CanvasState) func(socketio.Conn,string)  {
	return func(s socketio.Conn, msg string) {
		// append canvas
		var data []int
		json.Unmarshal([]byte(msg), &data)
		state.Canvas = append(state.Canvas, data)
		
		s.Emit("draw", data)	
	}
}

func OnError(s socketio.Conn, err error) {
	log.Println("error:", err)
}

func OnDisconnect(s socketio.Conn, reason string) {
	log.Println("disconnect", reason)
}