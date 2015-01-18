package bin

import (
	"os"

	"github.com/pengux/conf"
)

var Config *conf.Conf

// LoadConfiguration loads configuration parameters with the following
// order: env, file, arguments and defaults
func LoadConfiguration(file string) {
	defaults := map[string]interface{}{
		"DB_NAME":     "neightbour_social",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_PORT":     5432,
		"DB_SSL_MODE": "disable",
		// BcryptCost is used to set the cost in BCrypt hashing
		"BCRYPT_COST": 12,
		// TokenLength is used to set the length of a token in bytes
		"TOKEN_LENGTH":        32,
		"PASSWORD_MIN_LENGTH": 8,
		// Connection timeout in seconds
		"HTTP_TIMEOUT": 3600,
		// Maximum number of database connections for the application, 0 = Infinity
		"DB_MAXCONN": 0,
		// Maximum number of idle database connections, 0 = Infinity or DB_MAXCONN
		"DB_MAXIDLE": 0,
	}

	loader := conf.NewLoader()

	loader.Env()

	if file != "" {
		if _, err := os.Stat(file); err == nil {
			loader.File(file)
		}
	}

	loader.Argv().Defaults(defaults)

	c, err := loader.Load()
	if err != nil {
		panic(err)
	}

	Config = c
}
