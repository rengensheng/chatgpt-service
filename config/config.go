package config

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hyahm/golog"
	"gopkg.in/yaml.v2"
)

type Database struct {
	Username        string `yaml:"username"`
	Pwd             string `yaml:"pwd"`
	DriverName      string `yaml:"driver_name"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	DatabaseName    string `yaml:"database_name"`
	TableSchemaName string `yaml:"table_schema_name"`
	ShowSQL         bool   `yaml:"show_sql"`
}

type Service struct {
	Port      string `yaml:"port"`
	UploadDir string `yaml:"upload_dir"`
}

type Log struct {
	LogFilePath string `yaml:"log_file_path"`
	LogFileName string `yaml:"log_file_name"`
	Day         int    `yaml:"day"`
}

type Config struct {
	OpenAIKey     string   `yaml:"open_ai_key"`
	Database      Database `yaml:"database"`
	Service       Service  `yaml:"service"`
	Log           Log      `yaml:"log"`
	SecretKey     string   `yaml:"secret_key"`
	OpenAIBaseUel string   `yaml:"openai_base_url"`
}

var config Config

func init() {
	var configPath = "./config.yaml"
	if len(os.Args) > 1 {
		configPath = os.Args[1]
	}
	config, err := LoadConfig(configPath)
	golog.Info("配置加载完成...", config)
	if err != nil {
		log.Panic("配置加载错误:", err.Error())
	}
}

func LoadConfig(filepath string) (*Config, error) {
	configFile, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	yaml.Unmarshal(configFile, &config)
	return &config, nil
}

func GetConConfig() *Config {
	return &config
}
