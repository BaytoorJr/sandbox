package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetCompanyinfoTable(ctx *context.Context) table.Table {

	companyInfo := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := companyInfo.GetInfo().HideFilterArea()

	info.SetTable("public.company_info").SetTitle("Companyinfo").SetDescription("Companyinfo")

	formList := companyInfo.GetForm()

	formList.SetTable("public.company_info").SetTitle("Companyinfo").SetDescription("Companyinfo")

	return companyInfo
}
