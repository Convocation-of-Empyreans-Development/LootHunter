package esi

import (
	"LootHunter/types"
	"github.com/antihax/goesi"
	"net/http"
	"sync"
)

var UserAgent = "MALRO/LootHunter, built by Squish Padecain"

func CreateClient() *goesi.APIClient {
	return goesi.NewAPIClient(&http.Client{}, UserAgent)
}

func IDResolver(client *goesi.APIClient, workqueue types.IDResolverChannels, wg sync.WaitGroup) {
	defer wg.Done()
	for {
		work := <-workqueue.IDs
		names, response, err := client.ESI.UniverseApi.PostUniverseNames(nil, []int32{work}, nil)
		if err != nil || response.StatusCode != 200 {
			panic(err)
		}
		workqueue.Names <- names[0].Name
	}
}
