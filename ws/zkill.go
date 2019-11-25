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

// Connect to the Zkillboard websocket API.
func CreateZkillWebsocket() websocket.Conn {
	connection, _, err := websocket.DefaultDialer.Dial(ZkillboardWebsocketUrl.String(), nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] websocket connected")
	return *connection
}

// Subscribe to the killstream meesage queue given a websocket with access to the Zkillboard API.
func SubscribeToKillfeed(ws websocket.Conn) {
	err := ws.WriteMessage(websocket.TextMessage, []byte(`{"action":"sub", "channel":"killstream"}`))
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] successfully subscribed to killfeed")
}

// Repeatedly reads messages from the Zkillboard websocket and unmarshals the raw JSON data into structures for
// further processing. Detects duplicate killmails which are sometimes sent in batches over the queue, excluding
// these from processing.
// Only intended for use as a goroutine.
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
