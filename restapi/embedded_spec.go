package restapi

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "swagger": "2.0",
  "info": {
    "title": "example",
    "version": "0.0.0"
  },
  "paths": {
    "/items": {
      "get": {
        "security": [
          {
            "OauthSecurity": [
              "read"
            ]
          }
        ],
        "responses": {
          "200": {
            "description": "items"
          }
        }
      }
    }
  },
  "securityDefinitions": {
    "OauthSecurity": {
      "type": "oauth2",
      "flow": "accessCode",
      "authorizationUrl": "https://accounts.google.com/o/oauth2/auth",
      "tokenUrl": "https://www.googleapis.com/oauth2/v4/token",
      "scopes": {
        "read": "read items"
      }
    }
  }
}`))
}
