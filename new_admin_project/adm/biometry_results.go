package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetBiometryresultsTable(ctx *context.Context) table.Table {

	biometryResults := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := biometryResults.GetInfo().HideFilterArea()

	info.SetTable("public.biometry_results").SetTitle("Biometryresults").SetDescription("Biometryresults")

	formList := biometryResults.GetForm()

	formList.SetTable("public.biometry_results").SetTitle("Biometryresults").SetDescription("Biometryresults")

	return biometryResults
}
