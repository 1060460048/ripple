package ripple

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Configs  debug, release
type Configs struct {
	DebugOn bool `json:"debug_on"`
	Debug   Config
	Release Config
}

// Config config
type Config struct {
	Domain       string `json:"domain"`
	Static       string `json:"static"`
	Templates    string `json:"templates"`
	SignupEnable bool   `json:"signup_enable"`
	Database     *DatabaseConfig
}

// DatabaseConfig the database configuration
type DatabaseConfig struct {
	Dialect  string `json:"dialect"`
	Host     string `json:"host"`
	DbName   string `json:"dbname"`
	User     string `json:"user"`
	Password string `json:"password"`
}

// Current loaded config
var config *Config

// initConfig load the config file for application
func initConfig(e *echo.Echo) {
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

// NewConfig initial configuration from the config.json file and return a config instance
func NewConfig(e *echo.Echo) *Config {
	initConfig(e)
	return config
}

// GetConfig return the configuration
func GetConfig() *Config {
	return config
}

// GetDbConfig return the db config instance
func GetDbConfig() *DatabaseConfig {
	return config.Database
}

// GetStaticPath return the static path string
func GetStaticPath() string {
	return Getwd(config.Static)
}

// GetTemplatePath return the template path string
func GetTemplatePath() string {
	return Getwd(config.Templates)
}

// Getwd return the path's abs path string
func Getwd(path string) string {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workingDir, path)
}

// GetDbWithGorm return the github.com/jinzhu/gorm db
func GetDbWithGorm() (*gorm.DB, error) {
	cfg := GetDbConfig()
	dialect := cfg.Dialect
	host := cfg.Host
	dbname := cfg.DbName
	user := cfg.User
	password := cfg.Password

	connURI := ""
	switch dialect {
	case "mysql":
		connURI = user + ":" + password + "@tcp(" + host + ":3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	default:
		dialect = "mysql"
		connURI = user + ":" + password + "@tcp(" + host + ":3306)/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	}

	Logger.Info(fmt.Sprintf("[gorm] db_dialect: %s", dialect))
	Logger.Info(fmt.Sprintf("[gorm] db_connURI: %s", connURI))

	db, err := gorm.Open(dialect, connURI)
	return db, err
}
