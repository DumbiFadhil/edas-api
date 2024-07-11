package models

type Alternative struct {
	Name   string             `json:"name"`
	Scores map[string]float64 `json:"scores"`
}

type Criteria struct {
	Name   string  `json:"name"`
	Weight float64 `json:"weight"`
	Type   string  `json:"type"` // Add the Type field
}

type EDASRequest struct {
	Alternatives []Alternative `json:"alternatives"`
	Criteria     []Criteria    `json:"criteria"`
}

type EDASResponse struct {
	Ranking []RankedAlternative `json:"ranking"`
}

type RankedAlternative struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}
