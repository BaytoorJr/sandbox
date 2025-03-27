package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetProjectsTable(ctx *context.Context) table.Table {

	projects := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := projects.GetInfo().HideFilterArea()

	info.SetTable("public.projects").SetTitle("Projects").SetDescription("Projects")

	formList := projects.GetForm()

	formList.SetTable("public.projects").SetTitle("Projects").SetDescription("Projects")

	return projects
}
