package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetLoanagreementsTable(ctx *context.Context) table.Table {

	loanAgreements := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := loanAgreements.GetInfo().HideFilterArea()

	info.SetTable("public.loan_agreements").SetTitle("Loanagreements").SetDescription("Loanagreements")

	formList := loanAgreements.GetForm()

	formList.SetTable("public.loan_agreements").SetTitle("Loanagreements").SetDescription("Loanagreements")

	return loanAgreements
}
