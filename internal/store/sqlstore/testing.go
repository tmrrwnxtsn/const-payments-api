package sqlstore

import (
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"strings"
	"testing"
)

// TestDB инициализирует тестовое подключение к базе данных и truncate-функцию для очищения тестовых таблиц.
func TestDB(t *testing.T, dsn string) (*sqlx.DB, func(...string)) {
	t.Helper()

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		t.Fatal(err)
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			truncateTablesQuery := fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))
			_, err = db.Exec(truncateTablesQuery)
			if err != nil {
				t.Fatal(err)
			}
		}

		if err = db.Close(); err != nil {
			t.Fatal(err)
		}
	}
}
