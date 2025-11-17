package config

type DatabaseConfig struct {
	DatabaseURL string
}

func NewDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		DatabaseURL: getEnv("DATABASE_URL"),
	}
}
