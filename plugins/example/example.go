package example

import (
	c "github.com/hongweikkx/qzy-admin/modules/config"
	"github.com/hongweikkx/qzy-admin/modules/service"
	"github.com/hongweikkx/qzy-admin/plugins"
)

type Example struct {
	*plugins.Base
}

func NewExample() *Example {
	return &Example{
		Base: &plugins.Base{PlugName: "example"},
	}
}

func (e *Example) InitPlugin(srv service.List) {
	e.InitBase(srv, "example")
	e.App = e.initRouter(c.Prefix(), srv)
}
