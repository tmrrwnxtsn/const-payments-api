package sqlstore_test

import (
	"github.com/tmrrwnxtsn/const-payments-api/internal/config"
	"os"
	"testing"
)

var dsn string

func TestMain(m *testing.M) {
	dsn = os.Getenv(config.EnvVariablesPrefix + "DSN_TEST")
	if dsn == "" {
		dsn = "postgres://127.0.0.1/const_payments_db_test?sslmode=disable&user=postgres&password=qwerty"
	}

	os.Exit(m.Run())
}
