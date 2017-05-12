package backend

import (
	"net/http"

	"github.com/goadesign/goa"

	"github.com/MiCHiLU/goapp-scaffold/app"
)

type itemController struct {
	*goa.Controller
}

func newItemsController(service *goa.Service) *itemController {
	return &itemController{Controller: service.NewController("itemController")}
}

func (c *itemController) Get(ctx *app.GetItemsContext) error {
	return ctx.OK(nil)
}

func handleFunc(
	w http.ResponseWriter,
	r *http.Request,
) {
	return
}
