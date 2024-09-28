package pkg_config

type config struct {
	AppName      string `env:"APP_NAME"`
	AppPort      string `env:"APP_PORT"`
	AppEnv       string `env:"APP_ENV"`
	AppLogPath   string `env:"APP_LOG_PATH"`
	AppJwtSecret string `env:"APP_JWT_SECRET"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDbName   string `env:"POSTGRES_DB_NAME"`
}
