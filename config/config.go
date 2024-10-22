package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Application Application
	Mysql       Mysql
	Redis       Redis
	Rabbitmq    Rabbitmq
	ZapLogger   ZapLogger
}

type Application struct {
	HTTPServer                      HTTPServer
	EnableProfiling                 bool
	ProfilingPort                   int
	GracefulShutdownTimeoutInSecond int
}

type HTTPServer struct {
	Port string
}

type Mysql struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       string
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Rabbitmq struct {
	User                        string
	Password                    string
	Host                        string
	Port                        string
	Vhost                       string
	ReconnectSecond             int
	BufferSize                  int
	MaxRetryPolicy              int
	ChannelCleanerTimerInSecond int
}

type ZapLogger struct {
	Filename   string `json:"filename"`
	LocalTime  string `json:"local_time"`
	MaxSize    string `json:"max_size"`
	MaxBackups string `json:"max_backups"`
	MaxAge     string `json:"max_age"`
}

var AppConfig Config

func LoadConfig() error {

	var cfg Config

	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	cfg.Application.HTTPServer.Port = os.Getenv("HTTP_PORT")

	cfg.ZapLogger.Filename = os.Getenv("LOGGER_file_name")
	cfg.ZapLogger.LocalTime = os.Getenv("LOGGER_local_time")
	cfg.ZapLogger.MaxSize = os.Getenv("LOGGER_max_size")
	cfg.ZapLogger.MaxBackups = os.Getenv("LOGGER_max_backups")
	cfg.ZapLogger.MaxAge = os.Getenv("LOGGER_max_age")

	cfg.Mysql.Username = os.Getenv("DB_USERNAME")
	cfg.Mysql.Password = os.Getenv("DB_PASSWORD")
	cfg.Mysql.Host = os.Getenv("DB_HOST")
	cfg.Mysql.DB = os.Getenv("DB_NAME")
	cfg.Mysql.Port = os.Getenv("DB_PORT")

	cfg.Rabbitmq.User = os.Getenv("RABBITMQ_USER")
	cfg.Rabbitmq.Password = os.Getenv("RABBITMQ_PASSWORD")
	cfg.Rabbitmq.Host = os.Getenv("RABBITMQ_HOST")
	cfg.Rabbitmq.Port = os.Getenv("RABBITMQ_PORT")

	AppConfig = cfg

	return nil

}
