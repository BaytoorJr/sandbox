package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymenttaxesTable(ctx *context.Context) table.Table {

	paymentTaxes := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := paymentTaxes.GetInfo().HideFilterArea()

	info.SetTable("public.payment_taxes").SetTitle("Paymenttaxes").SetDescription("Paymenttaxes")

	formList := paymentTaxes.GetForm()

	formList.SetTable("public.payment_taxes").SetTitle("Paymenttaxes").SetDescription("Paymenttaxes")

	return paymentTaxes
}
