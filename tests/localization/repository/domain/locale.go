package domain

type Localization struct {
	ID           string                 `bson:"_id"`
	Translations map[string]interface{} `bson:"translations"`
}
