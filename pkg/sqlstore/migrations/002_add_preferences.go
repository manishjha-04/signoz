package migrations

import (
	"context"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
	"go.signoz.io/signoz/pkg/sqlstore"
)

type addPreferences struct {
	config sqlstore.MigrationConfig
}

func NewAddPreferencesMigrationFactory() sqlstore.MigrationFactory {
	return sqlstore.NewMigrationFactory(newAddPreferences)
}

func newAddPreferences(config sqlstore.MigrationConfig) sqlstore.Migration {
	return &addPreferences{config: config}
}

func (migration *addPreferences) Register(migrations *migrate.Migrations) error {
	if err := migrations.Register(migration.Up, migration.Down); err != nil {
		return err
	}

	return nil
}

func (migration *addPreferences) Up(ctx context.Context, db *bun.DB) error {
	// table:user_preference
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS user_preference (
		preference_id TEXT NOT NULL, 
		preference_value TEXT, 
		user_id TEXT NOT NULL, 
		PRIMARY KEY (preference_id,user_id), 
		FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE CASCADE ON DELETE CASCADE
	)`); err != nil {
		return err
	}

	// table:org_preference
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS org_preference (
		preference_id TEXT NOT NULL,
		preference_value TEXT,
		org_id TEXT NOT NULL,
		PRIMARY KEY (preference_id,org_id),
		FOREIGN KEY (org_id) REFERENCES organizations(id) ON UPDATE CASCADE ON DELETE CASCADE
	);`); err != nil {
		return err
	}

	return nil
}

func (migration *addPreferences) Down(ctx context.Context, db *bun.DB) error {
	return nil
}
