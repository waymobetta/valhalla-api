package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("godkjent", func() {
	BasePath("/v1/godkjent")

	Action("vis", func() {
		Description("vis godkjent")
		Routing(GET(""))
	})
	Response(OK, GodkjentListeMedia)
	Response(InternalServerError, StandardErrorMedia)
	Response(NotFound, StandardErrorMedia)

	Action("leggeTil", func() {
		Description("Legge til en godjkent")
		Routing(POST(""))
		Payload(LeggeTilGodkjentNyttelast)
		Response(OK, GodkjentMedia)
		Response(InternalServerError, StandardErrorMedia)
		Response(NotFound, StandardErrorMedia)
	})
})
