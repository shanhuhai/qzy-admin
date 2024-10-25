package example

import (
	"github.com/hongweikkx/qzy-admin/context"
	"github.com/hongweikkx/qzy-admin/modules/auth"
	"github.com/hongweikkx/qzy-admin/modules/db"
	"github.com/hongweikkx/qzy-admin/modules/service"
)

func (e *Example) initRouter(prefix string, srv service.List) *context.App {

	app := context.NewApp()
	route := app.Group(prefix)
	route.GET("/example", auth.Middleware(db.GetConnection(srv)), e.TestHandler)

	return app
}
