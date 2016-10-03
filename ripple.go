package ripple

import (
	"fmt"
	"github.com/bmbstack/ripple/middleware/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
	"os"
)

var Logger *logger.Logger
var baseRipple *Ripple
var firstRegController bool = true
var firstRegModel bool = true
var line1 string = "=============================="
var line2 string = "================================"

// Ripple ripple struct
type Ripple struct {
	Echo   *echo.Echo
	Config *Config
	Model  *Model
}

func init() {
	Logger = NewLogger()
	baseRipple = NewRipple()

	// Initial the DB with "github.com/jinzhu/gorm"
	baseModel := baseRipple.Model
	if baseModel != nil && !baseModel.IsOpenDB() {
		err := baseModel.OpenWithConfig()
		if err != nil {
			Logger.Error(err.Error())
			panic(err)
		}
	} else {
		model, err := NewModelWithConfig()
		if err != nil {
			Logger.Error(err.Error())
			panic(err)
		}
		baseRipple.Set(model)
	}
}

// NewRipple new a ripple instance
func NewRipple() *Ripple {
	localEcho := echo.New()

	r := &Ripple{}
	r.Set(localEcho)
	r.Set(NewConfig(localEcho))
	r.Set(NewModel())

	r.Echo.Use(mw.Recover())
	r.Echo.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}, rx_bytes=${rx_bytes}, tx_bytes=${tx_bytes}\n",
	}))

	// Set render
	r.Echo.SetRenderer(NewRenderer())
	r.Echo.Static("/static", GetStaticPath())

	return r
}

func NewLogger() *logger.Logger {
	log, err := logger.NewLogger("ripple", 1, os.Stdout)

	if err != nil {
		log.Error(err.Error())
		panic(err) // Check for error
	}
	return log
}

// Set set the ripple value
func (baseRipple *Ripple) Set(value interface{}) {
	switch value.(type) {
	case *echo.Echo:
		baseRipple.Echo = value.(*echo.Echo)
	case *Config:
		baseRipple.Config = value.(*Config)
	case *Model:
		baseRipple.Model = value.(*Model)
	}
}

// GetModel  return ripple model
func GetModel() *Model {
	return baseRipple.Model
}

// RegisterControllers register a controller for ripple App
func RegisterController(c Controller) {
	if firstRegController {
		fmt.Println(fmt.Sprintf("%s%s%s",
			color.White(line1),
			color.Bold(color.Green("Controller information")),
			color.White(line1)))
	}
	AddController(*baseRipple.Echo, c)
	firstRegController = false
}

// RegisterModels registers models in the global ripple App.
func RegisterModels(models ...interface{}) {
	if firstRegModel {
		fmt.Println(fmt.Sprintf("%s%s%s",
			color.White(line2),
			color.Bold(color.Green("Model information")),
			color.White(line2)))
	}
	baseRipple.Model.AddModels(models...)
	firstRegModel = false
}

// Migrate runs migrations on the global ripple App.
func Migrate() {
	baseRipple.Model.AutoMigrateAll()
}

// Run run ripple application
func Run() {
	Migrate()
	Logger.Info(fmt.Sprintf("Ripple ListenAndServe: %s", color.Green(baseRipple.Config.Domain)))
	baseRipple.Echo.Run(standard.New(baseRipple.Config.Domain))
}
