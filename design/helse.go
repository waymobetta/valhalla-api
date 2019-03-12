package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

var _ = Resource("helse", func() { // Resources group related API endpoints
	BasePath("/v1/helse") // together. They map to REST resources for REST

	NoSecurity()

	DefaultMedia(HelseMedia) // services.

	Action("vis", func() { // Actions define a single API endpoint together
		Description("OK hvis sunn")
		Routing(GET(""))
		Response(OK, HelseMedia)
		Response(InternalServerError, StandardErrorMedia)
		Response(NotFound, StandardErrorMedia)
	})
})
