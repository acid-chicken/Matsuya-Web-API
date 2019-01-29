package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("v1", func() {
	BasePath("v1")
	Action("random", func() {
		Routing(GET("random"))
		Description("ランダムで松屋のメニュー名を返す")
		Response(OK, ArrayOf(String), "application/json")
		Response(InternalServerError)
	})
})

var _ = Resource("v2", func() {
	BasePath("v2")
	DefaultMedia(V2Media)
	Action("random", func() {
		Routing(GET("random"))
		Description("ランダムで松屋のメニュー名を返す")
		Response(OK)
		Response(InternalServerError)
	})
})

var _ = Resource("v3", func() {
	BasePath("v3")
	DefaultMedia(V3Media)
	Action("random", func() {
		Routing(GET("random"))
		Description("ランダムで松屋のメニュー名を返す")
		Response(OK)
		Response(InternalServerError)
	})
})

var _ = Resource("v4", func() {
	BasePath("v4")
	Action("random", func() {
		Routing(GET("random"))
		Description("ランダムで松屋のメニュー名を返す")
		Params(func() {
			Param("type", String, "メニューの種類")
			Param("name", String, "メニュー名")
		})
		Response(OK, V4Media)
	})
	Action("search", func() {
		Routing(GET("search"))
		Description("ランダムで松屋のメニュー名を返す")
		Params(func() {
			Param("type", String, "メニューの種類")
			Param("name", String, "メニュー名")
		})
		Response(OK, ArrayOf(V4Media))
		Response(InternalServerError)
	})
})
