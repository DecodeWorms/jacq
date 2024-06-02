package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

const (
	appEnv         = "APP_ENV"
	servicePort    = "USER_SERVICE_PORT"
	autoReload     = "AUTO_RELOAD"
	databaseName   = "DATABASE_NAME"
	databaseURL    = "DATABASE_URL"
	serviceName    = "SERVICE_NAME"
	redisAddress   = "REDIS_ADDRESS"
	redisPassword  = "REDIS_PASSWORD"
	redisDb        = "REDIS_DB"
	secretKey      = "SECRET_KEY"
	gomailHostName = "GOMAI_NAME"
	gomailPort     = "GOMAIL_PORT"
	gmailPassword  = "GMAIL_PASSWORD"
	accountSid     = "ACCOUNT_SID"
	authToken      = "AUTH_TOKEN"
)

type source interface {
	GetEnv(key string, fallback string) string
	GetEnvBool(key string, fallback bool) bool
	GetEnvInt(key string, fallback int) int
}

type OSSource struct {
	source //nolint
}

func (o OSSource) GetEnv(key string, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func (o OSSource) GetEnvBool(key string, fallback bool) bool {
	b := o.GetEnv(key, "")
	if len(b) == 0 {
		return fallback
	}
	v, err := strconv.ParseBool(b)
	if err != nil {
		return fallback
	}
	return v
}

func (o OSSource) GetEnvInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		result, err := strconv.Atoi(value)
		if err != nil {
			return fallback
		}
		return result
	}
	return fallback
}

type Config struct {
	AppEnv        string
	ServicePort   string
	DatabaseName  string
	DatabaseURL   string
	AutoReload    bool
	ServiceName   string
	RedisAddress  string
	RedisPassword string
	RedisDb       int
	SecretKey     string
	GomailName    string
	GomailPort    int
	GmailPassword string
	AccountSid    string
	AuthToken     string
}

func ImportConfig(source source) Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appEnv := source.GetEnv(appEnv, "")
	port := source.GetEnv(servicePort, "8001")
	autoReload := source.GetEnvBool(autoReload, false)
	databaseName := source.GetEnv(databaseName, "Jacq")
	databaseURL := source.GetEnv(databaseURL, "mongodb://127.0.0.1:27017")
	serviceName := source.GetEnv(serviceName, "jacq")
	redisAddress := source.GetEnv(redisAddress, "localhost:6379")
	redisPassword := source.GetEnv(redisPassword, "")
	redisDb := source.GetEnvInt(redisDb, 0)
	secretKey := source.GetEnv(secretKey, "jacq-jwt-secret-key")
	gomailName := source.GetEnv(gomailHostName, "smtp.gmail.com")
	gomailPort := source.GetEnvInt(gomailPort, 587)
	gmailPass := source.GetEnv(gmailPassword, "fywtgbzlwqksolrr")
	accountSid := source.GetEnv(accountSid, "AC38ad1aaba69dbec742dc553854e04041")
	authToken := source.GetEnv(authToken, "00c4e9d61c68ccc10d600ebd6d99e5d7")

	return Config{
		AppEnv:        appEnv,
		ServicePort:   port,
		AutoReload:    autoReload,
		DatabaseName:  databaseName,
		DatabaseURL:   databaseURL,
		ServiceName:   serviceName,
		RedisAddress:  redisAddress,
		RedisPassword: redisPassword,
		RedisDb:       redisDb,
		SecretKey:     secretKey,
		GomailName:    gomailName,
		GomailPort:    gomailPort,
		GmailPassword: gmailPass,
		AccountSid:    accountSid,
		AuthToken:     authToken,
	}
}
