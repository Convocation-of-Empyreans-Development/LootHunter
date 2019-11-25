package ws

import (
	"LootHunter/types"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
	_ "time"
)

var ZkillboardWebsocketUrl = url.URL{Scheme: "wss", Host: "zkillboard.com:2096"}
var MaxPayloadSize = 1048576

func CreateZkillWebsocket() websocket.Conn {
	connection, _, err := websocket.DefaultDialer.Dial(ZkillboardWebsocketUrl.String(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] websocket connected")
	return *connection
}

func SubscribeToKillfeed(ws websocket.Conn) {
	err := ws.WriteMessage(websocket.TextMessage, []byte(`{"action":"sub", "channel":"killstream"}`))
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] successfully subscribed to killfeed")
}

func ZkillWebsocketLoop(ws websocket.Conn, feed chan types.Killmail, wg sync.WaitGroup) {
	defer wg.Done()

	var data = make([]byte, MaxPayloadSize)
	var latestKill types.Killmail
	var hashesSeen = make(map[string]bool)
	for {
		_, packet, err := ws.NextReader()
		if err != nil {
			panic(err)
		}
		bytesRead, _ := packet.Read(data)
		err = json.Unmarshal(data[:bytesRead], &latestKill)
		if err != nil {
			panic(err)
		}
		if _, found := hashesSeen[latestKill.Zkb.Hash]; !found {
			hashesSeen[latestKill.Zkb.Hash] = true
		} else {
			fmt.Printf("[?] recv dup with hash %s\n", latestKill.Zkb.Hash)
			continue
		}
		feed <- latestKill
	}
}
