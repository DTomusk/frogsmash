package config

type MessageConfig struct {
	RedisAddress  string
	RedisUsername string
	RedisPassword string
	StreamName    string
	GroupName     string
	ConsumerID    string
}

func NewMessageConfig() *MessageConfig {
	return &MessageConfig{
		RedisAddress:  getEnv("REDIS_ADDRESS"),
		RedisUsername: getEnv("REDIS_USERNAME"),
		RedisPassword: getEnv("REDIS_PASSWORD"),
		StreamName:    getEnv("STREAM_NAME"),
		GroupName:     getEnv("GROUP_NAME"),
		ConsumerID:    getEnv("CONSUMER_ID"),
	}
}
