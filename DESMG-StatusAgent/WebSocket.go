package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

var wsConn *websocket.Conn

func initWebSocket(u url.URL) {
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	wsConn = conn
	if err != nil {
		log.Println(err.Error())
	}
}

func reconnectWebSocket(u url.URL) {
	initWebSocket(u)
	for {
		if wsConn == nil {
			log.Println("Reconnection failed, retrying...")
			time.Sleep(time.Second)
			initWebSocket(u)
		}
	}
}

func send(deviceId string, data MonitorData) {
	if wsConn == nil {
		return
	}
	message, _ := json.Marshal(JsonData{
		DeviceID: deviceId,
		Data:     data,
	})
	err := wsConn.WriteMessage(websocket.BinaryMessage, message)
	if err != nil {
		log.Println(err.Error())
	}
}

func closeWebSocket() {
	if wsConn != nil {
		err := wsConn.Close()
		if err != nil {
			return
		}
	}
}
