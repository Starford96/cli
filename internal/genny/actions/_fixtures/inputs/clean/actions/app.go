package actions

import (
	"sync"

	"github.com/gobuffalo/buffalo"
)

var (
	app     *buffalo.App
	appOnce sync.Once
)

func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{})
		app.GET("/", HomeHandler)

		app.ServeFiles("/", assetsBox) // serve files from the public directory
	})

	return app
}
