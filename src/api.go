package main

import (
	"net/http"

	"appengine/user"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

const (
	clientId = "***.apps.googleusercontent.com"
)

var (
	scopes    = []string{endpoints.EmailScope}
	clientIds = []string{clientId, endpoints.APIExplorerClientID}
	audiences = []string{clientId}
)

func RegisterService() (rpcService *endpoints.RPCService, err error) {
	api := &Api{}
	rpcService, err = endpoints.RegisterService(api,
		"api", "v1", "api", true)
	if err != nil {
		return
	}

	info := rpcService.MethodByName("Items").Info()
	info.Name, info.HTTPMethod, info.Path, info.Desc =
		"item.list", "GET", "items", "List items."
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIds, audiences

	return
}

type Api struct {
}

func (api *Api) Items(r *http.Request,
	req *ItemsRequestMessage, resp *ItemsResponseMessage) (err error) {

	c := endpoints.NewContext(r)
	_, err = getCurrentUser(c)
	if err != nil {
		return
	}

	return
}

func getCurrentUser(c endpoints.Context) (u *user.User, err error) {
	u, err = endpoints.CurrentUser(c, scopes, audiences, clientIds)
	if err != nil {
		c.Infof("%v", err)
		err = endpoints.UnauthorizedError
		return
	}
	if u == nil {
		err = endpoints.UnauthorizedError
		return
	}
	return
}
