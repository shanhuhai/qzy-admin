package page

import (
	"bytes"

	"github.com/hongweikkx/qzy-admin/context"
	"github.com/hongweikkx/qzy-admin/modules/config"
	"github.com/hongweikkx/qzy-admin/modules/db"
	"github.com/hongweikkx/qzy-admin/modules/logger"
	"github.com/hongweikkx/qzy-admin/modules/menu"
	"github.com/hongweikkx/qzy-admin/plugins/admin/models"
	"github.com/hongweikkx/qzy-admin/template"
	"github.com/hongweikkx/qzy-admin/template/types"
)

// SetPageContent set and return the panel of page content.
func SetPageContent(ctx *context.Context, user models.UserModel, c func(ctx interface{}) (types.Panel, error), conn db.Connection) {

	panel, err := c(ctx)

	if err != nil {
		logger.ErrorCtx(ctx, "SetPageContent %+v", err)
		panel = template.WarningPanel(ctx, err.Error())
	}

	tmpl, tmplName := template.Get(ctx, config.GetTheme()).GetTemplate(ctx.IsPjax())

	ctx.AddHeader("Content-Type", "text/html; charset=utf-8")

	buf := new(bytes.Buffer)

	err = tmpl.ExecuteTemplate(buf, tmplName, types.NewPage(ctx, &types.NewPageParam{
		User:         user,
		Menu:         menu.GetGlobalMenu(user, conn, ctx.Lang()).SetActiveClass(config.URLRemovePrefix(ctx.Path())),
		Panel:        panel.GetContent(config.IsProductionEnvironment()),
		Assets:       template.GetComponentAssetImportHTML(ctx),
		TmplHeadHTML: template.Default(ctx).GetHeadHTML(),
		TmplFootJS:   template.Default(ctx).GetFootJS(),
		Iframe:       ctx.IsIframe(),
	}))
	if err != nil {
		logger.ErrorCtx(ctx, "SetPageContent %+v", err)
	}
	ctx.WriteString(buf.String())
}
