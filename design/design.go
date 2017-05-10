package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("example", func() {
	Title("example")
})

var OAuth2 = OAuth2Security("OAuth2", func() {
	AccessCodeFlow(
		"https://accounts.google.com/o/oauth2/auth",  //authorizationURL
		"https://www.googleapis.com/oauth2/v4/token", //tokenURL
	)
	Scope("read", "read")
})

var _ = Resource("items", func() {
	BasePath("/items")
	Action("get", func() {
		Routing(GET(""))
		Response(OK)
	})
	Security(OAuth2, func() {
		Scope("read")
	})
})
