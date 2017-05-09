package app

import (
	"net/http"
)

func init() {
	http.HandleFunc("/", handleFunc)
}
