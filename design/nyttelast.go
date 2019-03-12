package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// LeggeTilGodkjentNyttelast ...
var LeggeTilGodkjentNyttelast = Type("LeggeTilGodkjentNyttelast", func() {
	Description("Legge til godkjent på listen")
	Attribute("navn", String, "Navn på godkjent")
	Attribute("adresse", String, "Ethereum adresse")
	Required(
		"navn",
		"adresse",
	)
})
