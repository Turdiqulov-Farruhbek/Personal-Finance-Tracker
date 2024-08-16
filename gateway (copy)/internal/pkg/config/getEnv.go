package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
	// "github.com/spf13/cast"
)

type Config struct {
  GRPCPort string

  PostgresHost     string
  PostgresPort     int
  PostgresUser     string
  PostgresPassword string
  PostgresDatabase string
  LogPath          string
  KafkaUrl         string
  MongoUrl         string
  NotificationUrl  string
  BudgetUrl        string
  AuthUrl          string
  RedisUrl         string


  DefaultOffset string
  DefaultLimit  string
}

func Load() Config {
  if err := godotenv.Load(); err != nil {
    fmt.Println("No .env file found")
  }

  config := Config{}

  config.GRPCPort = cast.ToString(getOrReturnDefaultValue("GRPC_PORT", ":"))

  config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localst"))
  config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 542))
  config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "pogres"))
  config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "1234"))
  config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "proycts"))
  config.LogPath = cast.ToString(getOrReturnDefaultValue("LOG_PATH", "logs/info.log"))
  config.KafkaUrl = cast.ToString(getOrReturnDefaultValue("KAFKA_URL", "q"))
  config.MongoUrl = cast.ToString(getOrReturnDefaultValue("MONGO_URL", "q"))
  config.NotificationUrl = cast.ToString(getOrReturnDefaultValue("NOTIFICATION_URL", "q"))
  config.BudgetUrl = cast.ToString(getOrReturnDefaultValue("BUDGET_URL", "q"))
  config.AuthUrl = cast.ToString(getOrReturnDefaultValue("AUTH_URL", "q"))
  config.RedisUrl = cast.ToString(getOrReturnDefaultValue("REDIS_URL", "q"))


  config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
  config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))


  return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
  val, exists := os.LookupEnv(key)

  if exists {
    return val
  }

  return defaultValue
}