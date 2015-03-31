package app

import (
	"net/http"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

func init() {
	if _, err := RegisterService(); err != nil {
		panic(err.Error())
	}
	endpoints.HandleHTTP()
	//http.HandleFunc("/", handleFunc)
}

func RegisterService() (rpcService *endpoints.RPCService, err error) {
	api := &Api{}
	rpcService, err = endpoints.RegisterService(api,
		"api", "v1", "api", true)
	if err != nil {
		return
	}
	return
}

type Api struct {
}
