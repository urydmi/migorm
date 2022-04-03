package migrations

import (
	"errors"

	"github.com/urydmi/migorm"
	"gorm.io/gorm"
)

func init() {
	migorm.RegisterMigration(&migrationFirst{})
}

type migrationFirst struct{}

func (m *migrationFirst) Up(db *gorm.DB, log migorm.Logger) error {

	err := errors.New("implement me")

	return err
}

func (m *migrationFirst) Down(db *gorm.DB, log migorm.Logger) error {

	err := errors.New("implement me")

	return err
}
