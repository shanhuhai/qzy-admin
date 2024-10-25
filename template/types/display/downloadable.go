package display

import (
	"html/template"

	"github.com/hongweikkx/qzy-admin/context"
	"github.com/hongweikkx/qzy-admin/template/types"
)

type Downloadable struct {
	types.BaseDisplayFnGenerator
}

func init() {
	types.RegisterDisplayFnGenerator("downloadable", new(Downloadable))
}

func (d *Downloadable) Get(ctx *context.Context, args ...interface{}) types.FieldFilterFn {
	return func(value types.FieldModel) interface{} {
		param := args[0].([]string)

		u := value.Value

		if len(param) > 0 {
			u = param[0] + u
		}

		return template.HTML(`
<a href="` + u + `" download="` + value.Value + `" target="_blank" class="text-muted">
	<i class="fa fa-download"></i> ` + value.Value + `
</a>
`)
	}
}
