package utils

import (
	"LootHunter/types"
	_ "LootHunter/ws"
	"errors"
	"sync"
)

// Processes and distills the raw killmail data into only the required elements for other use in the application.
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

// Takes a list of items from the killmail and returns the ISK value of the dropped items.
func GetLootValue(items []types.ZkbItem) string {
	panic(errors.New("GetLootValue is not implemented"))
}

// Resolves the item ID of the ship lost in the killmail into the name of the ship.
func ResolveShip(id int) string {
	panic(errors.New("ResolveShip is not implemented"))
}

// Resolves the system ID from which the killmail was generated to the name of the system.
func ResolveSystemID(systemID int) string {
	panic(errors.New("ResolveSystemID is not implemented"))
}
