package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetDevicesTable(ctx *context.Context) table.Table {

	devices := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := devices.GetInfo().HideFilterArea()

	info.SetTable("public.devices").SetTitle("Devices").SetDescription("Devices")

	formList := devices.GetForm()

	formList.SetTable("public.devices").SetTitle("Devices").SetDescription("Devices")

	return devices
}
