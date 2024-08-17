package types

import (
	"fmt"
	"net/http"
)

type ControllerInterface interface {
	Init()
}

type Controller struct {
	BasePath string
	Routes   []Route
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
