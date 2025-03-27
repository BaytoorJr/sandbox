package adm

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
)

func GetSchemamigrationsTable(ctx *context.Context) table.Table {

	schemaMigrations := table.NewDefaultTable(ctx, table.DefaultConfigWithDriver("postgresql"))

	info := schemaMigrations.GetInfo().HideFilterArea()

	info.SetTable("public.schema_migrations").SetTitle("Schemamigrations").SetDescription("Schemamigrations")

	formList := schemaMigrations.GetForm()

	formList.SetTable("public.schema_migrations").SetTitle("Schemamigrations").SetDescription("Schemamigrations")

	return schemaMigrations
}
