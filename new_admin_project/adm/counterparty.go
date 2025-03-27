package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCounterpartyTable(ctx *context.Context) table.Table {

	counterparty := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := counterparty.GetInfo().HideFilterArea()

	info.SetTable("public.counterparty").SetTitle("Counterparty").SetDescription("Counterparty")

	formList := counterparty.GetForm()

	formList.SetTable("public.counterparty").SetTitle("Counterparty").SetDescription("Counterparty")

	return counterparty
}
