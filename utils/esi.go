package utils

import (
	"LootHunter/types"
)

func ResolveItemID(itemID int32, workqueue types.IDResolverChannels) string {
	workqueue.IDs <- itemID
	return <-workqueue.Names
}
