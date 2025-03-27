package repository

import (
	"context"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type LocaleRepository interface {
	GetLocalizations(ctx context.Context, bundle *i18n.Bundle) error
}
