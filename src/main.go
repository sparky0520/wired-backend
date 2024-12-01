package main

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	// authMiddleware := middleware.AuthMiddleware()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		return nil
	})

}
