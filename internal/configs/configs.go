package configs

import (
	"os"

	"github.com/mitchellh/mapstructure"
	"github.com/moeru-ai/inventory/internal/meta"
	"github.com/samber/lo"
	"github.com/spf13/viper"
)

type APIServer struct {
	GrpcServerAddr string `json:"grpc_server_addr" yaml:"grpc_server_addr"`
	HTTPServerAddr string `json:"http_server_addr" yaml:"http_server_addr"`
}

type Tracing struct {
	OtelCollectorHTTP bool `json:"otel_collector_http" yaml:"otel_collector_http"`
	OtelStdoutEnabled bool `json:"otel_stdout_enabled" yaml:"otel_stdout_enabled"`
}

type Config struct {
	meta.Meta `json:"-" yaml:"-"`

	Env       string    `json:"env" yaml:"env"`
	Tracing   Tracing   `json:"tracing" yaml:"tracing"`
	APIServer APIServer `json:"api_server" yaml:"api_server"`
}

func defaultConfig() Config {
	return Config{
		Tracing: Tracing{
			OtelCollectorHTTP: false,
			OtelStdoutEnabled: false,
		},
		APIServer: APIServer{
			GrpcServerAddr: ":7922",
			HTTPServerAddr: ":7921",
		},
	}
}

func NewConfig(namespace string, app string, configFilePath string, envFilePath string) func() (*Config, error) {
	return func() (*Config, error) {
		configPath := getConfigFilePath(configFilePath)

		lo.Must0(viper.BindEnv("env"))

		lo.Must0(viper.BindEnv("tracing.otel_collector_http"))
		lo.Must0(viper.BindEnv("tracing.otel_stdout_enabled"))

		lo.Must0(viper.BindEnv("api_server.grpc_server_bind"))
		lo.Must0(viper.BindEnv("api_server.http_server_bind"))

		err := loadEnvConfig(envFilePath)
		if err != nil {
			return nil, err
		}

		err = readConfig(configPath)
		if err != nil {
			return nil, err
		}

		config := defaultConfig()

		err = viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
			c.TagName = "yaml"
		})
		if err != nil {
			return nil, err
		}

		meta.Env = config.Env
		if meta.Env == "" {
			meta.Env = os.Getenv("ENV")
		}

		config.Env = meta.Env
		config.App = app
		config.Namespace = namespace

		return &config, nil
	}
}

func NewTestConfig(envFilePath string) (*Config, error) {
	configPath := tryToMatchConfigPathForUnitTest("")

	if envFilePath != "" {
		err := loadEnvConfig("")
		if err != nil {
			return nil, err
		}
	}

	err := readConfig(configPath)
	if err != nil {
		return nil, err
	}

	config := defaultConfig()
	config.Env = "test"

	err = viper.Unmarshal(&config, func(c *mapstructure.DecoderConfig) {
		c.TagName = "yaml"
	})
	if err != nil {
		return nil, err
	}

	return &config, nil
}
