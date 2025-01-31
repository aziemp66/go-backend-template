package pkg_config

type config struct {
	AppName      string `env:"APP_NAME"`
	AppPort      string `env:"APP_PORT"`
	AppEnv       string `env:"APP_ENV"`
	AppLogPath   string `env:"APP_LOG_PATH"`
	AppJwtSecret string `env:"APP_JWT_SECRET"`

	MailFrontEndUrl string `env:"MAIL_FRONT_END_URL"`
	MailHost        string `env:"MAIL_HOST"`
	MailPort        int    `env:"MAIL_PORT"`
	MailUser        string `env:"MAIL_HOST"`
	MailPassword    string `env:"MAIL_PASSWORD"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresPort     string `env:"POSTGRES_PORT"`
	PostgresDbName   string `env:"POSTGRES_DB_NAME"`
}
