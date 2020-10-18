package config

import (
	"github.com/spf13/viper"
	"time"
)

const (
	defaultHttpPort               = "8000"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
	defaultServiceName            = "users-service"
	defaultLoggerLevel            = 5 // debug level for logrus
	defaultDbSSLMode              = "disable"
)

type (
	Config struct {
		DB           DBConfig
		HTTP         HTTPConfig
		LoggerLevel  int
		ServiceName  string
		PasswordSalt string
	}

	DBConfig struct {
		Host     string
		Port     string
		Name     string
		Username string
		Password string
		SSLMode  string
	}

	HTTPConfig struct {
		Port               string
		ReadTimeout        time.Duration
		WriteTimeout       time.Duration
		MaxHeaderMegabytes int
	}
)

func Init() (Config, error) {
	if err := setUpViper(); err != nil {
		return Config{}, err
	}

	return setConfig(), nil
}

func setConfig() Config {
	return Config{
		LoggerLevel:  viper.GetInt("logger.level"),
		ServiceName:  viper.GetString("service.name"),
		PasswordSalt: viper.GetString("password.salt"),
		HTTP: HTTPConfig{
			Port:               viper.GetString("http.port"),
			MaxHeaderMegabytes: viper.GetInt("http.max_header_megabytes"),
			ReadTimeout:        viper.GetDuration("http.timeouts.read"),
			WriteTimeout:       viper.GetDuration("http.timeouts.write"),
		},
		DB: DBConfig{
			Host:     viper.GetString("host"),
			Port:     viper.GetString("port"),
			Name:     viper.GetString("name"),
			Username: viper.GetString("user"),
			Password: viper.GetString("pass"),
			SSLMode:  viper.GetString("sslmode"),
		},
	}
}

func setUpViper() error {
	populateDefaults()

	if err := parseConfigFile(); err != nil {
		return err
	}

	return parseEnvVariables()
}

func parseEnvVariables() error {
	viper.SetEnvPrefix("db")
	if err := viper.BindEnv("host"); err != nil {
		return err
	}

	if err := viper.BindEnv("port"); err != nil {
		return err
	}

	if err := viper.BindEnv("name"); err != nil {
		return err
	}

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	if err := viper.BindEnv("pass"); err != nil {
		return err
	}

	return viper.BindEnv("sslmode")
}

func parseConfigFile() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}

func populateDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.max_header_megabytes", defaultHttpMaxHeaderMegabytes)
	viper.SetDefault("http.timeouts.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeouts.write", defaultHttpRWTimeout)

	viper.SetDefault("service.name", defaultServiceName)

	viper.SetDefault("logger.level", defaultLoggerLevel)

	viper.SetDefault("db.sslmode", defaultDbSSLMode)
}
