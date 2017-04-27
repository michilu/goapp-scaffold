package main

import (
	"net/http"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	if _, err := RegisterService(); err != nil {
		panic(err.Error())
	}
	endpoints.HandleHTTP()
	http.HandleFunc("/", handleFunc)
}
