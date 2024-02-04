package compendit

import (
	// "compendit/internal/infrastructure"
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	App      *AppConfig
	Database *DatabaseConfig
	Redis    *RedisConfig
	Storage  *StorageConfig
	Bugsnag  *BugsnagConfig
	Planne   *PlanneConfig
	// Gateways *infrastructure.GatewaysConfig
}

type AppConfig struct {
	Env   string `env:"APP_ENV" envDefault:"development"`
	Port  int    `env:"APP_PORT" envDefault:"8000"`
	Debug bool   `env:"APP_DEBUG" envDefault:"false"`
}

type DatabaseConfig struct {
	Host         string `env:"DB_HOST"`
	Port         int    `env:"DB_PORT" envDefault:"3306"`
	User         string `env:"DB_USERNAME"`
	Password     string `env:"DB_PASSWORD"`
	DatabaseName string `env:"DB_DATABASE"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     int    `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD"`
	UseTls   bool   `env:"REDIS_USE_TLS" envDefault:"false"`
}

type PlanneConfig struct {
	ForceHttpsRedirect   bool   `env:"PLANNE_FORCE_HTTPS_REDIRECT" envDefault:"false"`
	SitePublicHostSuffix string `env:"PLANNE_SITE_PUBLIC_HOST_SUFFIX"`
}

type StorageConfig struct {
	PublicBucketUrl string `env:"STORAGE_PUBLIC_BUCKET_URL"`
}

type BugsnagConfig struct {
	ApiKey string `env:"BUGSNAG_API_KEY"`
}

func (cfg Config) AppEnv() string {
	return cfg.App.Env
}

func (cfg Config) IsDevelopment() bool {
	return cfg.App.Env == "development"
}

func (cfg Config) PublicBucketUrl() string {
	return cfg.Storage.PublicBucketUrl
}

// func (cfg Config) GatewaysConfig() infrastructure.GatewaysConfig {
// 	return *cfg.Gateways
// }

func RetrieveConfig() (Config, error) {
	cfg := Config{
		&AppConfig{},
		&DatabaseConfig{},
		&RedisConfig{},
		&StorageConfig{},
		&BugsnagConfig{},
		&PlanneConfig{},
		// &infrastructure.GatewaysConfig{
		// 	false,
		// 	&infrastructure.CieloConfig{},
		// 	&infrastructure.BraspagConfig{},
		// 	&infrastructure.ZoopConfig{},
		// 	&infrastructure.HopyPayConfig{},
		// 	&infrastructure.ClearSaleConfig{},
		// },
	}

	if err := env.Parse(&cfg); err != nil {
		return cfg, fmt.Errorf("error while retrieving config: %w", err)
	}

	return cfg, nil
}
