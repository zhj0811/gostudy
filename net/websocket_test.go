package net

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/zhj0811/gostudy/define"
)

func HandleWebsocketEvent(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Websocket starts.")

	var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	wsConn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		fmt.Printf("Error in websocket: %s", err.Error())
		return
	}
	defer func() {
		fmt.Println("Websocket quits.")
		wsConn.Close()
	}()

	if err := HandleEvent(wsConn); err != nil {
		fmt.Printf("Error in websocket: %s", err.Error())
	}
}

func HandleEvent(wsConn *websocket.Conn) error {
	fmt.Println("Service HandleWebsocketEvent")
	defer func() {
		fmt.Println("Service HandleWebsocketEvent end")
	}()
	for {
		data := define.Data{Key: "123", Value: "sdga"}
		resultJSON, _ := data.MarshalJSON()
		if err := wsConn.WriteMessage(websocket.TextMessage, resultJSON); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
}

func TestWebsocket(t *testing.T) {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	http.HandleFunc("/event/websocket", HandleWebsocketEvent)
	go func() {
		var err error
		err = http.ListenAndServe(":8080", nil)
		t.Log("server listening")
		if err != nil {
			t.Log(err.Error())
			s <- syscall.SIGTERM
		}
	}()
	<-s
	t.Log("Websocket server Exits.")
}
