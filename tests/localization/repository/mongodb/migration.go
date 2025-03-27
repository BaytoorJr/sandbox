package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"golangProject.com/localization/repository/domain"
	"log"
	"time"
)

func (s *Store) migrate() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := s.client.Database(dbName).Collection("localizations")

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return fmt.Errorf("ошибка проверки наличия данных: %w", err)
	}

	if count == 0 {
		// Английские переводы
		enTranslations := map[string]interface{}{
			"BuyingCookies": map[string]interface{}{
				"one":   "You're buying 1 cookie.",
				"other": "You're buying {{.PluralCount}} cookies.",
				"zero":  "You did not buy anything",
				"many":  "You're buying {{.PluralCount}} cookies.",
				"few":   "few",
			},
			"BuyingApples": map[string]interface{}{
				"one":   "You are buying 1 apple",
				"other": "You are buying {{.PluralCount}} apples: other",
				"zero":  "You did not buy any apple",
				"many":  "You are buying {{.PluralCount}} apples",
				"few":   "You are buying few apples",
			},
		}

		// Русские переводы
		ruTranslations := map[string]interface{}{
			"BuyingCookies": map[string]interface{}{
				"one":   "ОДИН",
				"zero":  "НОЛЬ",
				"many":  "Ты покупаешь {{.PluralForm}} печенье",
				"other": "ДРУГОЕ",
				"few":   "Немножко",
			},
			"BuyingApples": map[string]interface{}{
				"one":   "Ты покупаешь 1 яблоко",
				"zero":  "Ты не купил яблок",
				"many":  "Ты купил {{.PluralCount}} яблок",
				"other": "Ты купил {{.PluralCount}} яблок: other",
				"few":   "Ты купил несколько яблок",
			},
		}

		defaultLocalizations := []interface{}{
			domain.Localization{
				ID:           "en",
				Translations: enTranslations,
			},
			domain.Localization{
				ID:           "ru",
				Translations: ruTranslations,
			},
		}

		_, err = collection.InsertMany(ctx, defaultLocalizations)
		if err != nil {
			return fmt.Errorf("ошибка вставки данных: %w", err)
		}

		log.Println("Локализации добавлены в MongoDB.")
	} else {
		log.Println("Локализации уже существуют, миграция пропущена.")
	}

	return nil
}
