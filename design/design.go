package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const (
	appID = "example"
)

var _ = API(appID, func() {
	Title(appID)
})

var _ = Resource("items", func() {
	BasePath("/items")
	Action("get", func() {
		Routing(GET(""))
		Response(OK)
	})
})
