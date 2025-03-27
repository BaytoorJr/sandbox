package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetRefreshtokensTable(ctx *context.Context) table.Table {

	refreshTokens := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := refreshTokens.GetInfo().HideFilterArea()

	info.SetTable("public.refresh_tokens").SetTitle("Refreshtokens").SetDescription("Refreshtokens")

	formList := refreshTokens.GetForm()

	formList.SetTable("public.refresh_tokens").SetTitle("Refreshtokens").SetDescription("Refreshtokens")

	return refreshTokens
}
