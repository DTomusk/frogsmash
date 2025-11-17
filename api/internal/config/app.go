package config

type AppConfig struct {
	AllowedOrigin       string
	AppURL              string
	KFactor             float64
	MaxFileSize         int64
	Port                string
	ScoreUpdateInterval int
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		AllowedOrigin:       getEnv("ALLOWED_ORIGIN"),
		AppURL:              getEnv("APP_URL"),
		KFactor:             getFloat("KFACTOR"),
		MaxFileSize:         getInt64("MAX_FILE_SIZE_MB") * 1024 * 1024, // Convert MB to bytes
		Port:                getEnv("PORT"),
		ScoreUpdateInterval: getInt("SCORE_UPDATE_INTERVAL_SECONDS"),
	}
}
