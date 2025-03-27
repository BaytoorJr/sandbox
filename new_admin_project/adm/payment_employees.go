package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetPaymentemployeesTable(ctx *context.Context) table.Table {

	paymentEmployees := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := paymentEmployees.GetInfo().HideFilterArea()

	info.SetTable("public.payment_employees").SetTitle("Paymentemployees").SetDescription("Paymentemployees")

	formList := paymentEmployees.GetForm()

	formList.SetTable("public.payment_employees").SetTitle("Paymentemployees").SetDescription("Paymentemployees")

	return paymentEmployees
}
