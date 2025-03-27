package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetProfilesTable(ctx *context.Context) table.Table {

	profiles := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := profiles.GetInfo().HideFilterArea()

	info.SetTable("public.profiles").SetTitle("Profiles").SetDescription("Profiles")

	formList := profiles.GetForm()

	formList.SetTable("public.profiles").SetTitle("Profiles").SetDescription("Profiles")

	return profiles
}
