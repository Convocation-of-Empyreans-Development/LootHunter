package types

import "time"

// A killmail as received from the Zkillboard API.
type Killmail struct {
	ID            int        `json:"killmail_id"`
	Time          time.Time  `json:"killmail_time"`
	SolarSystemID int        `json:"solar_system_id"`
	Zkb           ZkbData    `json:"zkb"`
	Victim        VictimData `json:"victim"`
}

// Zkillboard's item representation on a killmail.
type ZkbItem struct {
	Flag              int `json:"flag"`
	ItemTypeID        int `json:"item_type_id"`
	QuantityDropped   int `json:"quantity_dropped"`
	QuantityDestroyed int `json:"quantity_destroyed"`
	Singleton         int `json:"singleton"`
}

// Zkillboard's data on the character who suffered the ship loss in the killmail.
type VictimData struct {
	CharacterID   int       `json:"character_id"`
	CorporationID int       `json:"corporation_id"`
	DamageTaken   int       `json:"damage_taken"`
	Items         []ZkbItem `json:"items"`
	ShipTypeID    int       `json:"ship_type_id"`
}

// Zkillboard's generic killmail-related metadata.
type ZkbData struct {
	LocationID  int     `json:"locationID"`
	Hash        string  `json:"hash"`
	FittedValue float64 `json:"fittedValue"`
	TotalValue  float64 `json:"totalValue"`
	Points      int     `json:"points"`
	NPC         bool    `json:"npc"`
	Solo        bool    `json:"solo"`
	Awox        bool    `json:"awox"`
	ESILink     string  `json:"esi"`
	ZkillLink   string  `json:"url"`
}

// A shortened, processed form of the Killmail struct with all irrelevant data removed.
type AbbreviatedKillmail struct {
	SystemName       string
	ShipDestroyed    string
	DroppedItemValue string
	Time             time.Time
}
