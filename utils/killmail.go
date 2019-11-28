package utils

import (
	"LootHunter/types"
	"fmt"
	"sync"
)

// Processes and distills the raw killmail data into only the required elements for other use in the application.
func ProcessKillmailFeed(feed chan types.Killmail, idResolveQueue types.IDResolverChannels, appraisalQueue types.AppraisalQueue, wg sync.WaitGroup) {
	defer wg.Done()
	for {
		killmail := <-feed
		shortKillmail := types.AbbreviatedKillmail{
			SystemName:       ResolveItemID(killmail.SolarSystemID, idResolveQueue),
			ShipDestroyed:    ResolveItemID(killmail.Victim.ShipTypeID, idResolveQueue),
			DroppedItemValue: GetLootValue(killmail.Victim.Items, idResolveQueue, appraisalQueue),
			Time:             killmail.Time,
		}
		fmt.Printf("%+v\n", shortKillmail)
	}
}

// Takes a list of items from the killmail and returns the ISK value of the dropped items.
func GetLootValue(items []types.ZkbItem, idResolveQueue types.IDResolverChannels, appraisalQueue types.AppraisalQueue) float64 {
	itemList := GenerateItemList(items, idResolveQueue)
	if itemList == "" {
		return 0
	}
	appraisalQueue.ItemLists <- itemList
	return (<-appraisalQueue.Appraisals).Data.Totals.Sell
}
