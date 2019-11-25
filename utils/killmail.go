package utils

import (
	"LootHunter/types"
	_ "LootHunter/ws"
	"errors"
	"sync"
)

func ProcessKillmailFeed(feed chan types.Killmail, wg sync.WaitGroup) {
	defer wg.Done()
	for {
		killmail := <-feed
		_ = types.AbbreviatedKillmail{
			SystemName:       ResolveSystemID(killmail.SolarSystemID),
			ShipDestroyed:    ResolveShip(killmail.Victim.ShipTypeID),
			DroppedItemValue: GetLootValue(killmail.Victim.Items),
			Time:             killmail.Time,
		}
	}
}

func GetLootValue(items []types.ZkbItem) string {
	panic(errors.New("GetLootValue is not implemented"))
}

func ResolveShip(id int) string {
	panic(errors.New("ResolveShip is not implemented"))
}

func ResolveSystemID(systemID int) string {
	panic(errors.New("ResolveSystemID is not implemented"))
}
