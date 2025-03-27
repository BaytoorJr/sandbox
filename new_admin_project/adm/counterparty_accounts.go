package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCounterpartyaccountsTable(ctx *context.Context) table.Table {

	counterpartyAccounts := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := counterpartyAccounts.GetInfo().HideFilterArea()

	info.SetTable("public.counterparty_accounts").SetTitle("Counterpartyaccounts").SetDescription("Counterpartyaccounts")

	formList := counterpartyAccounts.GetForm()

	formList.SetTable("public.counterparty_accounts").SetTitle("Counterpartyaccounts").SetDescription("Counterpartyaccounts")

	return counterpartyAccounts
}
