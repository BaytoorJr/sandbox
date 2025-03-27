package main

import (
	"log"
	"newadmin.com/adm"
	"os"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"                 // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/postgres" // Import the sql driver
	_ "github.com/GoAdminGroup/themes/adminlte"                      // Import the theme

	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	// GoAdmin global configuration, can also be imported as a json file.
	cfg := config.Config{
		Databases: make(config.DatabaseList),
		UrlPrefix: "admin", // The url prefix of the website.
		// Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
		Language: language.EN,
	}

	cfg.Databases.Add("default", config.Database{
		Host:         "127.0.0.1",
		Port:         "5432",
		User:         "yelnur",
		Pwd:          "BTR2002Big",
		Name:         "globerce",
		MaxIdleConns: 50,
		MaxOpenConns: 150,
		Driver:       config.DriverPostgresql,
	},
	)

	// Add configuration and plugins, use the Use method to mount to the web framework.
	err := eng.AddConfig(&cfg).
		AddGenerators(adm.Generators).
		Use(r)
	if err != nil {
		log.Println(err.Error())
		os.Exit(-1)
	}

	_ = r.Run(":9033")
}
