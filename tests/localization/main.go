package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kit/log/level"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	libmongo "gitlab.globerce.com/freedom-business/libs/shared-libs/databases/mongo"
	liblogger "gitlab.globerce.com/freedom-business/libs/shared-libs/utils/logger"
	"golang.org/x/text/language"
	"golangProject.com/localization/repository/mongodb"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	var count int
	var lang string

	flag.IntVar(&count, "count", 0, "number of items to buy")
	flag.StringVar(&lang, "lang", "ru", "language to use")
	flag.Parse()

	// init structured logger for the service
	logger := liblogger.NewServiceLogger("directory-api")
	_ = level.Info(logger).Log("msg", "service started")

	mongoClient, err := libmongo.InitConnect(ctx,
		"localhost",
		"27017",
		"",
		"",
		"test",
	)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}

	store, err := mongodb.NewStore(mongoClient)
	if err != nil {
		panic(err)
	}

	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	err = store.Locale().GetLocalizations(ctx, bundle)
	if err != nil {
		_ = level.Error(logger).Log("exit", err)
		os.Exit(-1)
	}
	fmt.Println(bundle.LanguageTags())

	lc := i18n.NewLocalizer(bundle, lang)

	msg, err := lc.Localize(&i18n.LocalizeConfig{
		MessageID: "BuyingApples",
		//PluralCount:  count,
		TemplateData: map[string]interface{}{"ApplesCount": count},
		DefaultMessage: &i18n.Message{
			ID: "",
		},
	})
	if err != nil {
		log.Fatalf("Ошибка получения перевода: %v", err)
	}

	fmt.Println(msg)

	//babel-api
}
