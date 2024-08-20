package types

import (
	"fmt"
	"net/http"
)

type Controller struct {
	BasePath string
	Routes   []Route
}

type ControllerInterface interface {
	Init()
}

func (c *Controller) Init(mux *http.ServeMux) {

	for _, route := range c.Routes {
		path := fmt.Sprintf("%s %s", route.Method, c.BasePath)

		if route.Path != "" {
			path += route.Path
		}

		mux.HandleFunc(path, route.Handler)
		fmt.Println("Route added: ", path)
	}
}
