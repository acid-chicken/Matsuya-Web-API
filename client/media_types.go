// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "Matsuya-Web-API": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/makotia/Matsuya-Web-API/design
// --out=$(GOPATH)/src/github.com/makotia/Matsuya-Web-API
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Version 4 (default view)
//
// Identifier: application/vnd.me.makotia.matsuya.random+json; view=default
type MeMakotiaMatsuyaRandom struct {
	// カロリー
	Calorie int `form:"calorie" json:"calorie" yaml:"calorie" xml:"calorie"`
	// 炭水化物
	Carbohydrate float64 `form:"carbohydrate" json:"carbohydrate" yaml:"carbohydrate" xml:"carbohydrate"`
	// 説明
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// 紹介画像
	ImageURL string `form:"imageURL" json:"imageURL" yaml:"imageURL" xml:"imageURL"`
	// 脂質
	Lipid float64 `form:"lipid" json:"lipid" yaml:"lipid" xml:"lipid"`
	// メニュー名
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 価格
	Price int `form:"price" json:"price" yaml:"price" xml:"price"`
	// タンパク質
	Protein float64 `form:"protein" json:"protein" yaml:"protein" xml:"protein"`
	// 食塩相当量
	SaltEquivalent float64 `form:"saltEquivalent" json:"saltEquivalent" yaml:"saltEquivalent" xml:"saltEquivalent"`
	// ナトリウム
	Sodium int `form:"sodium" json:"sodium" yaml:"sodium" xml:"sodium"`
	// 種類
	Type string `form:"type" json:"type" yaml:"type" xml:"type"`
}

// Validate validates the MeMakotiaMatsuyaRandom media type instance.
func (mt *MeMakotiaMatsuyaRandom) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Type == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "type"))
	}

	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.ImageURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "imageURL"))
	}
	return
}

// DecodeMeMakotiaMatsuyaRandom decodes the MeMakotiaMatsuyaRandom instance encoded in resp body.
func (c *Client) DecodeMeMakotiaMatsuyaRandom(resp *http.Response) (*MeMakotiaMatsuyaRandom, error) {
	var decoded MeMakotiaMatsuyaRandom
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Version 2 (default view)
//
// Identifier: application/vnd.me.makotia.matsuya.v2+json; view=default
type MeMakotiaMatsuyaV2 struct {
	// メニュー名
	Menu string `form:"menu" json:"menu" yaml:"menu" xml:"menu"`
}

// Validate validates the MeMakotiaMatsuyaV2 media type instance.
func (mt *MeMakotiaMatsuyaV2) Validate() (err error) {
	if mt.Menu == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "menu"))
	}
	return
}

// DecodeMeMakotiaMatsuyaV2 decodes the MeMakotiaMatsuyaV2 instance encoded in resp body.
func (c *Client) DecodeMeMakotiaMatsuyaV2(resp *http.Response) (*MeMakotiaMatsuyaV2, error) {
	var decoded MeMakotiaMatsuyaV2
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// Version 3 (default view)
//
// Identifier: application/vnd.me.makotia.matsuya.v3+json; view=default
type MeMakotiaMatsuyaV3 struct {
	// カロリー
	Calorie int `form:"calorie" json:"calorie" yaml:"calorie" xml:"calorie"`
	// 炭水化物
	Carbohydrate float64 `form:"carbohydrate" json:"carbohydrate" yaml:"carbohydrate" xml:"carbohydrate"`
	// 説明
	Description string `form:"description" json:"description" yaml:"description" xml:"description"`
	// 紹介画像
	ImageURL string `form:"imageURL" json:"imageURL" yaml:"imageURL" xml:"imageURL"`
	// 脂質
	Lipid float64 `form:"lipid" json:"lipid" yaml:"lipid" xml:"lipid"`
	// メニュー名
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
	// 価格
	Price int `form:"price" json:"price" yaml:"price" xml:"price"`
	// タンパク質
	Protein float64 `form:"protein" json:"protein" yaml:"protein" xml:"protein"`
	// 食塩相当量
	SaltEquivalent float64 `form:"saltEquivalent" json:"saltEquivalent" yaml:"saltEquivalent" xml:"saltEquivalent"`
	// ナトリウム
	Sodium int `form:"sodium" json:"sodium" yaml:"sodium" xml:"sodium"`
}

// Validate validates the MeMakotiaMatsuyaV3 media type instance.
func (mt *MeMakotiaMatsuyaV3) Validate() (err error) {
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	if mt.Description == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "description"))
	}
	if mt.ImageURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "imageURL"))
	}
	return
}

// DecodeMeMakotiaMatsuyaV3 decodes the MeMakotiaMatsuyaV3 instance encoded in resp body.
func (c *Client) DecodeMeMakotiaMatsuyaV3(resp *http.Response) (*MeMakotiaMatsuyaV3, error) {
	var decoded MeMakotiaMatsuyaV3
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
