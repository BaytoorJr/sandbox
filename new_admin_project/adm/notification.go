package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetNotificationTable(ctx *context.Context) table.Table {

	notification := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := notification.GetInfo().HideFilterArea()

	info.SetTable("public.notification").SetTitle("Notification").SetDescription("Notification")

	formList := notification.GetForm()

	formList.SetTable("public.notification").SetTitle("Notification").SetDescription("Notification")

	return notification
}
