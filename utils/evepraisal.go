package utils

import (
	"LootHunter/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func GenerateItemList(items []types.ZkbItem, workqueue types.IDResolverChannels) string {
	var itemList []string
	for _, item := range items {
		if item.QuantityDropped == 0 {
			continue
		}
		// append returns a copy of the slice
		itemList = append(itemList, types.EvepraisalItem{
			Name:     ResolveItemID(item.ItemTypeID, workqueue),
			Quantity: item.QuantityDropped,
		}.String())
	}
	if len(itemList) == 0 {
		return ""
	}
	return strings.Join(itemList, "\n")
}

func AppraisalQueue(queue types.AppraisalQueue, config types.EvepraisalServerData, wg *sync.WaitGroup) {
	defer wg.Done()
	apiUrl := fmt.Sprintf("http://%s/appraisal.json?market=%s&persist=%s", config.Url, config.Market, config.Persist)
	for {
		itemList := <-queue.ItemLists
		queue.Appraisals <- GetAppraisal(apiUrl, itemList)
	}
}

func GetAppraisal(apiUrl string, itemList string) types.Appraisal {
	var data []byte
	var appraisal types.Appraisal
	response, err := http.Post(apiUrl, "text/plain", bytes.NewBufferString(itemList))
	if err != nil || response.StatusCode != 200 {
		panic(err)
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &appraisal)
	if err != nil {
		panic(err)
	}
	return appraisal
}
