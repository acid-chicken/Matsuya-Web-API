package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var V2Media = MediaType("application/vnd.me.makotia.matsuya.v2+json", func() {
	Description("Version 2")
	Attributes(func() {
		Attribute("menu", String, "メニュー名")
		Required("menu")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("menu")
	})
})

var V3Media = MediaType("application/vnd.me.makotia.matsuya.v3+json", func() {
	Description("Version 3")
	Attributes(func() {
		Attribute("name", String, "メニュー名")
		Attribute("price", Integer, "価格")
		Attribute("calorie", Integer, "カロリー")
		Attribute("protein", Number, "タンパク質")
		Attribute("lipid", Number, "脂質")
		Attribute("carbohydrate", Number, "炭水化物")
		Attribute("sodium", Integer, "ナトリウム")
		Attribute("saltEquivalent", Number, "食塩相当量")
		Attribute("description", String, "説明")
		Attribute("imageURL", String, "紹介画像")
		Required(
			"name",
			"price",
			"calorie",
			"protein",
			"lipid",
			"carbohydrate",
			"sodium",
			"saltEquivalent",
			"description",
			"imageURL",
		)
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("name")
		Attribute("price")
		Attribute("calorie")
		Attribute("protein")
		Attribute("lipid")
		Attribute("carbohydrate")
		Attribute("sodium")
		Attribute("saltEquivalent")
		Attribute("description")
		Attribute("imageURL")
	})
})

var V4Media = MediaType("application/vnd.me.makotia.matsuya.v4+json", func() {
	Description("Version 4")
	Attributes(func() {
		Attribute("name", String, "メニュー名")
		Attribute("type", String, "種類")
		Attribute("price", Integer, "価格")
		Attribute("calorie", Integer, "カロリー")
		Attribute("protein", Number, "タンパク質")
		Attribute("lipid", Number, "脂質")
		Attribute("carbohydrate", Number, "炭水化物")
		Attribute("sodium", Integer, "ナトリウム")
		Attribute("saltEquivalent", Number, "食塩相当量")
		Attribute("description", String, "説明")
		Attribute("imageURL", String, "紹介画像")
		Required(
			"name",
			"type",
			"price",
			"calorie",
			"protein",
			"lipid",
			"carbohydrate",
			"sodium",
			"saltEquivalent",
			"description",
			"imageURL",
		)
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("name")
		Attribute("type")
		Attribute("price")
		Attribute("calorie")
		Attribute("protein")
		Attribute("lipid")
		Attribute("carbohydrate")
		Attribute("sodium")
		Attribute("saltEquivalent")
		Attribute("description")
		Attribute("imageURL")
	})
})
