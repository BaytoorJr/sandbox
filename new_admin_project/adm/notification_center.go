package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetNotificationcenterTable(ctx *context.Context) table.Table {

	notificationCenter := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := notificationCenter.GetInfo().HideFilterArea()

	info.SetTable("public.notification_center").SetTitle("Notificationcenter").SetDescription("Notificationcenter")

	formList := notificationCenter.GetForm()

	formList.SetTable("public.notification_center").SetTitle("Notificationcenter").SetDescription("Notificationcenter")

	return notificationCenter
}
