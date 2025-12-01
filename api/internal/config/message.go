package config

type MessageConfig struct {
	RedisAddress string
}

func NewMessageConfig() *MessageConfig {
	return &MessageConfig{
		RedisAddress: getEnv("REDIS_ADDRESS"),
	}
}
