package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

const (
	appID            = "example"
	google_audiences = "CLIENT-ID"

	authorizationURL = "https://accounts.google.com/o/oauth2/auth"
	google_issuer    = "https://accounts.google.com"
	google_jwks_uri  = "https://www.googleapis.com/oauth2/v3/certs"
)

var _ = API(appID, func() {
	Title(appID)
	Origin("http://localhost:8080", func() {
		Headers("Content-Type")
		Methods("GET")
	})
})

var JWT = OAuth2Security("google_id_token", func() {
	ImplicitFlow(authorizationURL)
	Scope("openid", "OpenID connect")
})

var _ = Resource("items", func() {
	BasePath("/items")
	Action("get", func() {
		Routing(GET(""))
		Response(OK)
	})
	Security(JWT, func() {
		Scope("openid")
		Metadata("swagger:extension:x-google-issuer", google_issuer)
		Metadata("swagger:extension:x-google-jwks_uri", google_jwks_uri)
		//Metadata("swagger:extension:x-google-audiences", google_audiences)
	})
})
