package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort string
	HTTPHost string

	PostgresHost            string
	PostgresPort            string
	PostgresUser            string
	PostgresPassword        string
	PostgresDatabase        string
	PostgresMaxConnections  string
	PostgresConnMaxIdleTime string // in minutes

	SecretKey string

	RedisHost       string
	RedisPort       string
	RedisDB         string
	RedisPassword   string
	RedisPoolSize   string
	RedisExpiryTime string

	AccessTokenName string
	AccessTokenTTL  string // in minutes

	RefreshTokenName string
	RefreshTokenTTL  string // in minutes

	Amazons3AccessKey string
	Amazons3SecretKey string
	Amazons3Region    string
	Amazons3Bucket    string
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	cnf := Config{}
	// http
	cnf.HTTPPort = getOrReturnDefaultValue("HTTP_PORT", 8080).(string)
	cnf.HTTPHost = getOrReturnDefaultValue("HTTP_HOST", "localhost").(string)
	//postgres
	cnf.PostgresHost = getOrReturnDefaultValue("POSTGRES_HOST", "localhost").(string)
	cnf.PostgresPort = getOrReturnDefaultValue("POSTGRES_PORT", 5432).(string)
	cnf.PostgresUser = getOrReturnDefaultValue("POSTGRES_USER", "postgres").(string)
	cnf.PostgresDatabase = getOrReturnDefaultValue("POSTGRES_DB", "postgres").(string)
	cnf.PostgresPassword = getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres").(string)
	cnf.PostgresMaxConnections = getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 10).(string)
	cnf.PostgresConnMaxIdleTime = getOrReturnDefaultValue("POSTGRES_CONN_MAX_IDLE_TIME", 10).(string)
	// jwt
	cnf.SecretKey = getOrReturnDefaultValue("SECRET_KEY0", "secret").(string)
	// redis
	cnf.RedisHost = getOrReturnDefaultValue("REDIS_HOST", "localhost").(string)
	cnf.RedisPort = getOrReturnDefaultValue("REDIS_PORT", "6379").(string)
	cnf.RedisDB = getOrReturnDefaultValue("REDIS_DB", "0").(string)
	cnf.RedisPassword = getOrReturnDefaultValue("REDIS_PASSWORD", "").(string)
	cnf.RedisPoolSize = getOrReturnDefaultValue("REDIS_POOL_SIZE", "10").(string)
	cnf.RedisExpiryTime = getOrReturnDefaultValue("REDIS_EXPIRY_TIME", "10").(string)
	// amazons3
	cnf.Amazons3AccessKey = getOrReturnDefaultValue("AWS_ACCESS_KEY_ID", "").(string)
	cnf.Amazons3SecretKey = getOrReturnDefaultValue("AWS_SECRET_ACCESS_KEY", "").(string)
	cnf.Amazons3Region = getOrReturnDefaultValue("AWS_REGION", "").(string)
	cnf.Amazons3Bucket = getOrReturnDefaultValue("AWS_BUCKET_NAME", "").(string)

	return &cnf
}
