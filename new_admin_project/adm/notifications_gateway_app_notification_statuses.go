package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetNotificationsgatewayappnotificationstatusesTable(ctx *context.Context) table.Table {

	notificationsGatewayAppNotificationStatuses := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := notificationsGatewayAppNotificationStatuses.GetInfo().HideFilterArea()

	info.SetTable("public.notifications_gateway_app_notification_statuses").SetTitle("Notificationsgatewayappnotificationstatuses").SetDescription("Notificationsgatewayappnotificationstatuses")

	formList := notificationsGatewayAppNotificationStatuses.GetForm()

	formList.SetTable("public.notifications_gateway_app_notification_statuses").SetTitle("Notificationsgatewayappnotificationstatuses").SetDescription("Notificationsgatewayappnotificationstatuses")

	return notificationsGatewayAppNotificationStatuses
}
