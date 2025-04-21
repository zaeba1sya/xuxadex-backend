package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type Config struct {
	App struct {
		Env string `yaml:"env" envconfig:"APP_ENV" required:"true"`
	} `yaml:"app"`
	Server struct {
		Host string `yaml:"host" envconfig:"APP_HOST" required:"true"`
		Port uint16 `yaml:"port" envconfig:"APP_PORT" required:"true"`
	} `yaml:"server"`
	DB struct {
		Dialect  string `yaml:"dialect" envconfig:"DB_DIALECT" required:"true"`
		Host     string `yaml:"host" envconfig:"DB_HOST" required:"true"`
		Port     uint16 `yaml:"port" envconfig:"DB_PORT" required:"true"`
		User     string `yaml:"user" envconfig:"DB_USER" required:"true"`
		Password string `yaml:"password" envconfig:"DB_PASSWORD" required:"true"`
		Name     string `yaml:"name" envconfig:"DB_NAME" required:"true"`
		SSL      bool   `yaml:"ssl" envconfig:"DB_SSL" required:"true"`
		Pool     struct {
			MaxOpenConns    int `yaml:"max_open_conns" envconfig:"DB_MAX_OPEN_CONNS" required:"true"`
			MaxIdleConns    int `yaml:"max_idle_conns" envconfig:"DB_MAX_IDLE_CONNS" required:"true"`
			ConnMaxLifetime int `yaml:"conn_max_lifetime" envconfig:"DB_CONN_MAX_LIFETIME" required:"true"`
		} `yaml:"pool" required:"true"`
	} `yaml:"db" required:"true"`
	Blockchain struct {
		NodeURL string `yaml:"node_url" envconfig:"BLOCKCHAIN_NODE_URL" required:"true"`
		// Temp field, change on contract list, when contract is deployed
		ContractAddress string `yaml:"contract_address" envconfig:"BLOCKCHAIN_CONTRACT_ADDRESS" required:"true"`
	} `yaml:"blockchain" required:"true"`
	Logger struct {
		Level    string `yaml:"level" envconfig:"LOG_LEVEL" required:"true"`
		Encoding string `yaml:"encoding" envconfig:"LOG_ENCODING" required:"true"`
	} `yaml:"logger" required:"true"`
}

func loadEnv(cfg *Config) error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	err := envconfig.Process("", cfg)
	if err != nil {
		return err
	}

	return nil
}

func loadYaml(cfg *Config) error {
	yamlLocation := flag.String("yaml", "./config.yaml", "Load config from file")

	f, err := os.Open(*yamlLocation)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return err
	}

	return nil
}

func InitConfig() *Config {
	flag.Parse()

	var cfg Config
	var err error

	cfgType := flag.String("cfgType", "env", "Choose config type")

	switch *cfgType {
	case "yaml":
		err = loadYaml(&cfg)
	default:
		err = loadEnv(&cfg)
	}

	if err != nil {
		panic(err)
	}

	return &cfg
}
