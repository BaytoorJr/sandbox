package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetBaseprofileTable(ctx *context.Context) table.Table {

	baseProfile := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := baseProfile.GetInfo().HideFilterArea()

	info.SetTable("public.base_profile").SetTitle("Baseprofile").SetDescription("Baseprofile")

	formList := baseProfile.GetForm()

	formList.SetTable("public.base_profile").SetTitle("Baseprofile").SetDescription("Baseprofile")

	return baseProfile
}
