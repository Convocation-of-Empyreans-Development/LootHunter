package main

import (
	"LootHunter/esi"
	"LootHunter/types"
	"LootHunter/utils"
	"LootHunter/ws"
	"flag"
	"sync"
)

var EvepraisalAddress = flag.String("evepraisal", "evepraisal.com", "URL for Evepraisal service")
var wg = &sync.WaitGroup{}

func main() {
	flag.Parse()

	sock := ws.CreateZkillWebsocket()
	ws.SubscribeToKillfeed(sock)

	feed := make(chan types.Killmail)
	wg.Add(1)
	go ws.ZkillWebsocketLoop(sock, feed, *wg)
	client := esi.CreateClient()
	idResolverWorkqueue := types.IDResolverChannels{
		IDs:   make(chan int32),
		Names: make(chan string),
	}
	appraisalWorkqueue := types.AppraisalQueue{
		ItemLists:  make(chan string),
		Appraisals: make(chan types.Appraisal),
	}
	wg.Add(1)
	go esi.IDResolver(client, idResolverWorkqueue, *wg)
	wg.Add(1)
	go utils.ProcessKillmailFeed(feed, idResolverWorkqueue, appraisalWorkqueue, *wg)
	wg.Add(1)
	go utils.AppraisalQueue(appraisalWorkqueue, types.EvepraisalServerData{
		Url:     *EvepraisalAddress,
		Market:  "jita",
		Persist: "no",
	}, wg)

	wg.Wait()
}
