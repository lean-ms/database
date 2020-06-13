package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-pg/pg"
)

type DbConfigTestItem struct {
	env      string
	database string
	user     string
}

var databaseVsEnvironmentTest = []DbConfigTestItem{
	{
		env:      "",
		database: "leanms_dabatase_package_development",
		user:     "postgres",
	},
	{
		env:      "development",
		database: "leanms_dabatase_package_development",
		user:     "postgres",
	},
	{
		env:      "test",
		database: "leanms_dabatase_package_test",
		user:     "postgres",
	},
	{
		env:      "production",
		database: "leanms_dabatase_package_production",
		user:     "USER_SET_IN_ENV_VARIABLE",
	},
}

func TestLoadingConfigFromEnv(t *testing.T) {
	for _, item := range databaseVsEnvironmentTest {
		setupEnvVars(item)
		config := LoadDatabaseConfig("../database.yml")
		if checkItemAndDatabaseOptionsMatches(item, config) {
			t.Error(getMismatchErrorMessage(item))
		}
	}
}

func setupEnvVars(item DbConfigTestItem) {
	os.Setenv("DATABASE_USER", "USER_SET_IN_ENV_VARIABLE")
	os.Setenv("LEANMS_ENV", item.env)
}

func checkItemAndDatabaseOptionsMatches(item DbConfigTestItem, config *pg.Options) bool {
	return item.database != config.Database || item.user != config.User
}

func getMismatchErrorMessage(item DbConfigTestItem) string {
	str := "Could not load database config for env %v"
	return fmt.Sprintf(str, item.env)
}
