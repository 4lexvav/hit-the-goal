package config

import cfg "github.com/Yalantis/go-config"

var config Config

func Get() *Config {
	return &config
}

func Load(filename string) error {
	return cfg.Init(&config, filename)
}

type (
	Config struct {
		AppName   string `json:"app_name"   envconfig:"APP_NAME"   default:"hit-the-goal"`
		ListenURL string `json:"listen_url" envconfig:"LISTEN_URL" default:":8080"`
		LogPreset string `json:"log_preset" envconfig:"LOG_PRESET" default:"development"`

		Postgres Postgres `json:"postgres"`
		Redis    Redis    `json:"redis"`
	}

	Postgres struct {
		Host         string       `json:"host"          envconfig:"POSTGRES_HOST"          default:"localhost"`
		Port         string       `json:"port"          envconfig:"POSTGRES_PORT"          default:"5432"`
		Database     string       `json:"database"      envconfig:"POSTGRES_DATABASE"      default:"goal"`
		User         string       `json:"user"          envconfig:"POSTGRES_USER"          default:"postgres"`
		Password     string       `json:"password"      envconfig:"POSTGRES_PASSWORD"      default:"12345"`
		PoolSize     int          `json:"pool_size"     envconfig:"POSTGRES_POOL_SIZE"     default:"10"`
		MaxRetries   int          `json:"max_retries"   envconfig:"POSTGRES_MAX_RETRIES"   default:"5"`
		ReadTimeout  cfg.Duration `json:"read_timeout"  envconfig:"POSTGRES_READ_TIMEOUT"  default:"10s"`
		WriteTimeout cfg.Duration `json:"write_timeout" envconfig:"POSTGRES_WRITE_TIMEOUT" default:"10s"`
	}

	Redis struct {
		Address  string `json:"address"   envconfig:"REDIS_ADDRESS"   default:"localhost:6379"`
		Password string `json:"password"  envconfig:"REDIS_PASSWORD"  default:"password"`
		PoolSize int    `json:"pool_size" envconfig:"REDIS_POOL_SIZE" default:"10"`
	}
)
