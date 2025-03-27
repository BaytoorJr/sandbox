package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetRolesTable(ctx *context.Context) table.Table {

	roles := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := roles.GetInfo().HideFilterArea()

	info.SetTable("public.roles").SetTitle("Roles").SetDescription("Roles")

	formList := roles.GetForm()

	formList.SetTable("public.roles").SetTitle("Roles").SetDescription("Roles")

	return roles
}
