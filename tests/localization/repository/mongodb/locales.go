package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"go.mongodb.org/mongo-driver/bson"
	"golangProject.com/localization/repository/domain"
	"log"
	"time"
)

type LocaleRepo struct {
	store *Store
}

func (r *LocaleRepo) GetLocalizations(ctx context.Context, bundle *i18n.Bundle) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := r.store.client.Database(dbName).Collection("localizations")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("ошибка получения переводов из MongoDB: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var loc domain.Localization
		if err := cursor.Decode(&loc); err != nil {
			return fmt.Errorf("ошибка декодирования локализации: %w", err)
		}

		// Преобразуем Translations в JSON
		data, err := json.Marshal(loc.Translations)
		if err != nil {
			return fmt.Errorf("ошибка сериализации переводов: %w", err)
		}

		// Создаем виртуальный путь, т.к. go-i18n использует путь для определения языка
		virtualPath := loc.ID + ".json"

		// Загружаем переводы прямо в bundle
		if _, err := bundle.ParseMessageFileBytes(data, virtualPath); err != nil {
			return fmt.Errorf("ошибка загрузки переводов для %s: %w", loc.ID, err)
		}

		log.Printf("Переводы для языка %s загружены", loc.ID)
	}

	if err := cursor.Err(); err != nil {
		return fmt.Errorf("ошибка при итерации по документам: %w", err)
	}

	return nil
}
