package app

import (
	"net/http"

	"appengine/user"
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
)

const (
	clientID = "***.apps.googleusercontent.com"
)

var (
	scopes    = []string{endpoints.EmailScope}
	clientIDs = []string{clientID, endpoints.APIExplorerClientID}
	audiences = []string{clientID}
)

func registerService() (rpcService *endpoints.RPCService, err error) {
	api := &API{}
	rpcService, err = endpoints.RegisterService(api,
		"api", "v1", "api", true)
	if err != nil {
		return
	}

	info := rpcService.MethodByName("Items").Info()
	info.Name, info.HTTPMethod, info.Path, info.Desc =
		"item.list", "GET", "items", "List items."
	info.Scopes, info.ClientIds, info.Audiences = scopes, clientIDs, audiences

	return
}

// API ...
type API struct {
}

// Items ...
func (api *API) Items(r *http.Request,
	req *ItemsRequestMessage, resp *ItemsResponseMessage) (err error) {

	c := endpoints.NewContext(r)
	_, err = getCurrentUser(c)
	if err != nil {
		return
	}

	return
}

func getCurrentUser(c endpoints.Context) (u *user.User, err error) {
	u, err = endpoints.CurrentUser(c, scopes, audiences, clientIDs)
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
