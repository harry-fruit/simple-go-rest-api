package types

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Controller struct {
	BasePath string
	Routes   []Route
}

type ControllerInterface interface {
	Init()
}

func (c *Controller) SetRoutes(mux *mux.Router) {

	for _, route := range c.Routes {
		path := c.BasePath + route.Path
		mux.Handle(path, http.HandlerFunc(route.Handler)).Methods(route.Method)
		fmt.Println("Route added: ", path)
	}
}
