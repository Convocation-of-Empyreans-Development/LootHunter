package types

import "time"

type Killmail struct {
	ID            int        `json:"killmail_id"`
	Time          time.Time  `json:"killmail_time"`
	SolarSystemID int        `json:"solar_system_id"`
	Zkb           ZkbData    `json:"zkb"`
	Victim        VictimData `json:"victim"`
}

type ZkbItem struct {
	Flag              int `json:"flag"`
	ItemTypeID        int `json:"item_type_id"`
	QuantityDropped   int `json:"quantity_dropped"`
	QuantityDestroyed int `json:"quantity_destroyed"`
	Singleton         int `json:"singleton"`
}

type VictimData struct {
	CharacterID   int       `json:"character_id"`
	CorporationID int       `json:"corporation_id"`
	DamageTaken   int       `json:"damage_taken"`
	Items         []ZkbItem `json:"items"`
	ShipTypeID    int       `json:"ship_type_id"`
}

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

type AbbreviatedKillmail struct {
	SystemName       string
	ShipDestroyed    string
	DroppedItemValue string
	Time             time.Time
}
