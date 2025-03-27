package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetOtpsTable(ctx *context.Context) table.Table {

	otps := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := otps.GetInfo().HideFilterArea()

	info.SetTable("public.otps").SetTitle("Otps").SetDescription("Otps")

	formList := otps.GetForm()

	formList.SetTable("public.otps").SetTitle("Otps").SetDescription("Otps")

	return otps
}
