package subjectField

type Model struct {
	ID           int     `json:"id,omitempty"`
	Rate1        uint32  `json:"rate_1"`
	Rate2        uint32  `json:"rate_2"`
	Rate3        uint32  `json:"rate_3"`
	Rate4        uint32  `json:"rate_4"`
	Rate5        uint32  `json:"rate_5"`
	Rate6        uint32  `json:"rate_6"`
	Rate7        uint32  `json:"rate_7"`
	Rate8        uint32  `json:"rate_8"`
	Rate9        uint32  `json:"rate_9"`
	Rate10       uint32  `json:"rate_10"`
	AverageScore float64 `json:"average_score"`
	Rank         uint32  `json:"rank,omitempty"`
	Year         uint32  `json:"year,omitempty"`
	Month        uint8   `json:"month,omitempty"`
	Date         uint8   `json:"date,omitempty"`
	Weekday      uint8   `json:"weekday,omitempty"`
}
