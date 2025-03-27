package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetKycresultsTable(ctx *context.Context) table.Table {

	kycResults := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := kycResults.GetInfo().HideFilterArea()

	info.SetTable("public.kyc_results").SetTitle("Kycresults").SetDescription("Kycresults")

	formList := kycResults.GetForm()

	formList.SetTable("public.kyc_results").SetTitle("Kycresults").SetDescription("Kycresults")

	return kycResults
}
