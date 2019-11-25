package main

import (
	"LootHunter/types"
	"LootHunter/utils"
	"LootHunter/ws"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	sock := ws.CreateZkillWebsocket()
	ws.SubscribeToKillfeed(sock)
	feed := make(chan types.Killmail)
	wg.Add(1)
	go ws.ZkillWebsocketLoop(sock, feed, *wg)
	wg.Add(1)
	go utils.ProcessKillmailFeed(feed, *wg)
	wg.Wait()
}
