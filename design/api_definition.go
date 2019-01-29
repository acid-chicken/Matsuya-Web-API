package design

import (
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Matsuya-Web-API", func() {
	Title("Matsuya-Web-API")
	Description("松屋のメニューをランダムで返すAPIとメニューのデータ(栄養値)を返すAPI")
	Host("matsuya.makotia.me")
	Scheme("https")
	Consumes("application/json")
	Produces("application/json")
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "swaggerui")
})
