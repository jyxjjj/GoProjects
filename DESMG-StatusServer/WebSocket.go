package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{}

func initWebSocketServer() {
	http.HandleFunc("/agent", handleMonitorData)
	addr := "127.0.0.1:7001"
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return
	}
}

func handleMonitorData(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for {
		_, message, err1 := conn.ReadMessage()
		if err1 != nil {
			log.Println(err.Error())
			return
		}
		log.Println(string(message))
		var jsonData JsonData
		err2 := json.Unmarshal(message, &jsonData)
		if err2 != nil {
			log.Println(err.Error())
			return
		}
		log.Println(jsonData.DeviceID)
	}
}
