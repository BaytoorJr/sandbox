package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetActionsTable(ctx *context.Context) table.Table {

	actions := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := actions.GetInfo().HideFilterArea()

	info.SetTable("public.actions").SetTitle("Actions").SetDescription("Actions")

	formList := actions.GetForm()

	formList.SetTable("public.actions").SetTitle("Actions").SetDescription("Actions")

	return actions
}
