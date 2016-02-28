package ripple

import (
	"encoding/json"
	"github.com/labstack/echo"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Configs struct {
	DebugOn bool `json:"debug_on"`
	Debug   Config
	Release Config
}

type Config struct {
	Domain       string `json:"domain"`
	Static       string `json:"static"`
	Templates    string `json:"templates`
	LogLevel     string `json:"log_level"`
	SignupEnable bool   `json:"signup_enable"`
	Database     DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Name     string
	User     string
	Password string
}

// Current loaded config
var config *Config

func LoadConfig(e *echo.Echo) {
	configFile := Getwd("config/config.json")
	var input = io.ReadCloser(os.Stdin)
	input, configErr := os.Open(configFile)
	if configErr != nil {
		panic(configErr)
	}
	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		panic(err)
	}

	configs := &Configs{}
	err = json.Unmarshal(jsonBytes, configs)
	if err != nil {
		panic(err)
	}

	debug := configs.DebugOn
	e.SetDebug(debug)
	if debug {
		config = &configs.Debug
	} else {
		config = &configs.Release
	}

	if !filepath.IsAbs(config.Static) {
		workingDir, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		config.Static = filepath.Join(workingDir, config.Static)
	}
}

func GetConfig() *Config {
	return config
}

func GetStaticPath() string {
	return Getwd(config.Static)
}

func GetTemplatePath() string {
	return Getwd(config.Templates)
}

func Getwd(path string) string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workingDir, path)
}
