package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetSwaggersTable(ctx *context.Context) table.Table {

	swaggers := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := swaggers.GetInfo().HideFilterArea()

	info.SetTable("public.swaggers").SetTitle("Swaggers").SetDescription("Swaggers")

	formList := swaggers.GetForm()

	formList.SetTable("public.swaggers").SetTitle("Swaggers").SetDescription("Swaggers")

	return swaggers
}
