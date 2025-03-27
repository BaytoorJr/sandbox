package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetProfileTable(ctx *context.Context) table.Table {

	profile := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := profile.GetInfo().HideFilterArea()

	info.SetTable("public.profile").SetTitle("Profile").SetDescription("Profile")

	formList := profile.GetForm()

	formList.SetTable("public.profile").SetTitle("Profile").SetDescription("Profile")

	return profile
}
