package types

import "fmt"

type EvepraisalItem struct {
	Name     string
	Quantity int
}

func (i EvepraisalItem) String() string {
	return fmt.Sprintf("%s x%d", i.Name, i.Quantity)
}

type AppraisalValues struct {
	Buy    float64 `json:"buy"`
	Sell   float64 `json:"sell"`
	Volume float64 `json:"volume"`
}

type AppraisalData struct {
	Totals AppraisalValues `json:"totals"`
}

type Appraisal struct {
	Data AppraisalData `json:"appraisal"`
}

type AppraisalQueue struct {
	ItemLists  chan string
	Appraisals chan Appraisal
}

type EvepraisalServerData struct {
	Url     string
	Market  string
	Persist string
}
