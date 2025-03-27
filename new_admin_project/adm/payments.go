package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymentsTable(ctx *context.Context) table.Table {

	payments := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := payments.GetInfo().HideFilterArea()

	info.SetTable("public.payments").SetTitle("Payments").SetDescription("Payments")

	formList := payments.GetForm()

	formList.SetTable("public.payments").SetTitle("Payments").SetDescription("Payments")

	return payments
}
