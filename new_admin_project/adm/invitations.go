package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetInvitationsTable(ctx *context.Context) table.Table {

	invitations := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := invitations.GetInfo().HideFilterArea()

	info.SetTable("public.invitations").SetTitle("Invitations").SetDescription("Invitations")

	formList := invitations.GetForm()

	formList.SetTable("public.invitations").SetTitle("Invitations").SetDescription("Invitations")

	return invitations
}
