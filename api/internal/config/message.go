package config

type MessageConfig struct {
	RedisAddress string
	StreamName   string
	GroupName    string
	ConsumerID   string
}

func NewMessageConfig() *MessageConfig {
	return &MessageConfig{
		RedisAddress: getEnv("REDIS_ADDRESS"),
		StreamName:   getEnv("STREAM_NAME"),
		GroupName:    getEnv("GROUP_NAME"),
		ConsumerID:   getEnv("CONSUMER_ID"),
	}
}
