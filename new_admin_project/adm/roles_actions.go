package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetRolesactionsTable(ctx *context.Context) table.Table {

	rolesActions := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := rolesActions.GetInfo().HideFilterArea()

	info.SetTable("public.roles_actions").SetTitle("Rolesactions").SetDescription("Rolesactions")

	formList := rolesActions.GetForm()

	formList.SetTable("public.roles_actions").SetTitle("Rolesactions").SetDescription("Rolesactions")

	return rolesActions
}
