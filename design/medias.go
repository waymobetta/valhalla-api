package design

import (
	. "github.com/goadesign/goa/design"        // Use . imports to enable the DSL
	. "github.com/goadesign/goa/design/apidsl" // Use . imports to enable the DSL
)

// HelseMedia ...
var HelseMedia = MediaType("application/vnd.helse+json", func() {
	Description("Helse")
	Attributes(func() { // Attributes define the media type shape.
		Attribute("status", String, "Status")
		Required("status")
	})
	View("default", func() { // View defines a rendering of the media type.
		Attribute("status")
	})
})

// GodkjentMedia ...
var GodkjentMedia = MediaType("application/vnd.godkjent+json", func() {
	Description("Vis godjkent")
	Attributes(func() {
		Attribute("navn", String, "Navn på godkjent")
		Attribute("adresse", String, "Ethereum adresse")
		Required(
			"navn",
			"adresse",
		)
	})
	View("default", func() {
		Attribute("navn")
		Attribute("adresse")
	})
})

// GodkjentListeMedia ...
var GodkjentListeMedia = MediaType("application/vnd.godkjentliste+json", func() {
	Description("Vis godkjent liste")
	Attributes(func() {
		Attribute("godkjentListe", CollectionOf(GodkjentMedia), "liste av godkjent")
		Required("godkjentListe")
	})
	View("default", func() {
		Attribute("godkjentListe")
	})
})
