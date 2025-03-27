package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetLicenseTable(ctx *context.Context) table.Table {

	license := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := license.GetInfo().HideFilterArea()

	info.SetTable("public.license").SetTitle("License").SetDescription("License")

	formList := license.GetForm()

	formList.SetTable("public.license").SetTitle("License").SetDescription("License")

	return license
}
