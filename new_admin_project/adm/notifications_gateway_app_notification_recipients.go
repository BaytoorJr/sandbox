package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetNotificationsgatewayappnotificationrecipientsTable(ctx *context.Context) table.Table {

	notificationsGatewayAppNotificationRecipients := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := notificationsGatewayAppNotificationRecipients.GetInfo().HideFilterArea()

	info.SetTable("public.notifications_gateway_app_notification_recipients").SetTitle("Notificationsgatewayappnotificationrecipients").SetDescription("Notificationsgatewayappnotificationrecipients")

	formList := notificationsGatewayAppNotificationRecipients.GetForm()

	formList.SetTable("public.notifications_gateway_app_notification_recipients").SetTitle("Notificationsgatewayappnotificationrecipients").SetDescription("Notificationsgatewayappnotificationrecipients")

	return notificationsGatewayAppNotificationRecipients
}
