package collection

type Info struct {
	Type       CollectionType `json:"type"`
	HasComment bool           `json:"hasComment"`
	Comment    string         `json:"comment"`
	Score      uint8          `json:"score"`
	AddTime    string         `json:"addTime"`
	EpStatus   uint8          `json:"epStatus"`
}
