package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"

	dspslds "github.com/adgear/dsp-lds/pkg/config"
)

const configParentDirectEnvVarName = "RTB_BIDDER_CONFIG_PARENT_DIRECT"
const envEnvVarName = "RTB_BIDDER_ENV"

type (
	// Config stores all configuration properties of the application.
	// The values are read by "viper" from a config file or environment variable.
	Config struct {
		App                          `yaml:"app"`
		Logger                       `yaml:"logger"`
		Metrics                      `yaml:"metrics"`
		Http                         `yaml:"http"`
		HealthCheck                  `yaml:"healthcheck"`
		dspslds.LocalDataStoreConfig `mapstructure:",squash"`
	}

	// App Application configuration settings.
	App struct {
		Name        string `yaml:"name"`        // Application name.
		Version     string `yaml:"version"`     // Application version in format "Major.Minor.Patch".
		Environment string `yaml:"environment"` // Execution environment.
	}

	// Logger configuration settings.
	Logger struct {
		Level  string `yaml:"level"`  // Logging level.
		Format string `yaml:"format"` // Logs format.
	}

	// Metrics Logger configuration settings.
	Metrics struct {
		Namespace string `yaml:"namespace"` // Metrics namespace. Required for groupping metrics from multiple services related to the project "neo"
		Subsystem string `yaml:"subsystem"` // Metrics subsystem, Required for groupping metrics from current service.
	}

	// Http Web server configuration settings.
	Http struct {
		Name                 string  `yaml:"name"`
		Port                 int     `yaml:"port"` // Server listening on port.
		MaxConnsPerIp        int     `yaml:"maxconnsperip"`
		MaxReqPerConn        int     `yaml:"maxreqperconn"`
		NotFoundSamplingRate float32 `yaml:"notfoundsamplingrate"`
		ReadTimeoutSeconds   int     `yaml:"readtimeoutseconds"`
		WriteTimeoutSeconds  int     `yaml:"writetimeoutseconds"`
		IdleTimeoutSeconds   int     `yaml:"idletimeoutseconds"`
	}

	// Health service configuration settings.
	HealthCheck struct {
		LoopIntervalSeconds int `yaml:"loopintervalseconds"` // Health checks loop interval
	}
)

var (
	// Variable to hold configuration.
	cfg     Config
	cfgOnce sync.Once
)

// ProvideConfig Return singleton of the configuration struct.
// Loads configuration from the file and override it with values from environment variables.
// If configuration already loaded, returns the configuration.
// Configuration read from the `{env}.yaml` file where `env“ is current environment loaded from `RTB_BIDDER_ENV` env variable.
// Default environment is "dev".
// Environment variables could be read from os environment variables(highest priority) and from .env files in following order:.
// - .env
// - .env.{env}
// - .env.{env}.local
func ProvideConfig() *Config {
	cfgOnce.Do(func() {
		var err error
		path := configPath()
		cfg, err = loadConfig(path)
		if err != nil {
			panic(err)
		}

	})
	return &cfg
}

func loadConfig(path string) (config Config, err error) {
	// The configuration name is the environment. The dufault environment is "dev".
	env := getEnv(envEnvVarName, "dev")
	viper.SetConfigName(env)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	loadEnv(env)

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

// Return the configuration directory. The configuration parent directory could be set by `RTB_BIDDER_CONFIG_PARENT_DIRECT` env variable.
// If env variable `RTB_BIDDER_CONFIG_PARENT_DIRECT` is not set, the parent configuration directory will be set to the current working directory.
func configPath() (path string) {
	workingDirectory, err := os.Getwd()

	if err != nil {
		panic(fmt.Errorf("failed to retrieve working directory: %w", err))
	}

	configParentDirectory := getEnv(configParentDirectEnvVarName, workingDirectory)

	if configParentDirectory, err = filepath.Abs(configParentDirectory); err != nil {
		panic(fmt.Errorf("failed to resolve CONFIG_PARENT_DIR: %w", err))
	}

	path = configParentDirectory + "/config"

	return
}

// The .env file is set more generic values. The .env.{env} where `env` is the current environment. for example .env.dev for example. Values from .env.{env} will override values.
// The .env.{env}.local will override values. The local file is for local debugging and is not under source control.
func loadEnv(env string) {
	_ = godotenv.Load(".env/.env." + env + ".local")
	_ = godotenv.Load(".env/.env." + env)
	_ = godotenv.Load(".env/.env")
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
