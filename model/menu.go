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
	Name           string  `name:"name" json:"name"`
	Type           string  `bson:"type" json:"type"`
	Price          int     `bson:"price" json:"price"`
	Calorie        int     `bson:"calorie" json:"calorie"`
	Protein        float64 `bson:"protein" json:"protein"`
	Lipid          float64 `bson:"lipid" json:"lipid"`
	Carbohydrate   float64 `bson:"carbohydrate" json:"carbohydrate"`
	Sodium         int     `bson:"sodium" json:"sodium"`
	SaltEquivalent float64 `bson:"saltEquivalent" json:"saltEquivalent"`
	Description    string  `bson:"description" json:"description"`
	ImageURL       string  `bson:"imageURL" json:"imageURL"`
}
