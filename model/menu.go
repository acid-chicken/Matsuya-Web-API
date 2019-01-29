package model

import "encoding/json"

func UnmarshalMatsuyaMenu(data []byte) (MatsuyaMenu, error) {
	var r MatsuyaMenu
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MatsuyaMenu) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MatsuyaMenu struct {
	Name           string  `json:"name"`
	Type           string  `json:"type"`
	Price          int64   `json:"price"`
	Calorie        int64   `json:"calorie"`
	Protein        float64 `json:"protein"`
	Lipid          float64 `json:"lipid"`
	Carbohydrate   int64   `json:"carbohydrate"`
	Sodium         int64   `json:"sodium"`
	SaltEquivalent float64 `json:"saltEquivalent"`
	Description    string  `json:"description"`
	ImageURL       string  `json:"imageURL"`
}
