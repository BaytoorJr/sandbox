package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetActiongroupsTable(ctx *context.Context) table.Table {

	actionGroups := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := actionGroups.GetInfo().HideFilterArea()

	info.SetTable("public.action_groups").SetTitle("Actiongroups").SetDescription("Actiongroups")

	formList := actionGroups.GetForm()

	formList.SetTable("public.action_groups").SetTitle("Actiongroups").SetDescription("Actiongroups")

	return actionGroups
}
