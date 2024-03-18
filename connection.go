package discord

import (
	"github.com/gorilla/websocket"
	"log"
	"os/signal"
	"os"
	"encoding/json"
	"fmt"
	"time"
	"runtime"
)

// # Connection Struct
type Connection struct {
	Socket		*websocket.Conn
	Shard 		int
	SessionId	string
	Sequence	int
	Timer		int
	Timestamp	int
	Buffers		int
	Heartbeats	int
}

// Payload Struct
type Payload struct {
	Opcode		int		`json:"op"`
	Sequence	int		`json:"s"`
	Type		string 	`json:"t"`
	Data		any		`json:"d"`
}


func identify(conn *websocket.Conn, token string) {
	identify_payload := Payload{
		Opcode: 2,
		Data: map[string]any{
			"token":   		token,
			"intents": 		Intent_ALL,
			"properties": 	map[string]any{
				"$os":      runtime.GOOS,
				"$browser": "github.com/grhx/disgord",
				"$device":  "golang",
			},
		},
	}
	json_msg, err := json.Marshal(identify_payload)
	if err != nil {
		log.Fatal(err)
	}
	conn.WriteMessage(websocket.TextMessage, json_msg)
}

func heartbeat(ch chan bool, conn *websocket.Conn) {
	defer close(ch)
	hb_msg, err := json.Marshal(Payload{Opcode: 1,Data:nil})
	if err != nil {
		log.Fatal(err)
	}
	conn.WriteMessage(websocket.TextMessage, hb_msg)
	for {
		time.Sleep(time.Second * 40)
		conn.WriteMessage(websocket.TextMessage, hb_msg)
	}
}

func (c *Client) dialGateway(ch chan bool, token string) {
	defer close(ch)

	println("Beginning of dial")

	// upgrader
	// var upgrader = websocket.Upgrader{
	// 	ReadBufferSize: 1024,
	// 	WriteBufferSize: 1024,
	// }
	// // handler
	// func handler(w http.http.ResponseWriter, r *http.Request) {

	// 	conn, err := upgrader.Upgrade(w, r, nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	return
	// }

	interrupt := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(interrupt, os.Interrupt)
	conn, resp, err := websocket.DefaultDialer.Dial("wss://gateway.discord.gg/?v=10&encoding=json", nil)
	fmt.Printf("GATEWAY CONNECTION RESPONSE: %v\n", resp)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer close(done)
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			var event Payload
			err = json.Unmarshal(message, &event)
			if err != nil {
				log.Fatal(err)
			}
			times := time.Now()
			log_debug := fmt.Sprintf("\x1b[42m%d/%d/%d %d:%d:%d LOG:\x1b[0m", times.Year(), times.Month(), times.Day(), times.Hour(), times.Minute(), times.Second())
			event_debug := fmt.Sprintf("%s %s", log_debug, fmt.Sprintf("op: %d, t: %s, s: %d", event.Opcode, event.Type, event.Sequence))
			fmt.Println(event_debug)
			switch Opcode(event.Opcode) {
			// dipatch
			case Opcode_DISPATCH:
			// heartbeat
			case Opcode_HEARTBEAT:
			// reconnect
			case Opcode_RECONNECT:
			// invalid session
			case Opcode_INVALID_SESSION:
			// hello
			case Opcode_HELLO:
				identify(conn, token)
				go heartbeat(ch, conn)
			// heartbeat acknowledge
			case Opcode_HEARTBEAT_ACK:
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
		}
	}

}