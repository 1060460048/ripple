package ripple

import (
	"fmt"
	"github.com/bmbstack/ripple/middleware/logger"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	mw "github.com/labstack/echo/middleware"
	"os"
)

var Logger *logger.Logger
var baseRipple *Ripple

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
			fmt.Println(err.Error())
		}
	} else {
		model, err := NewModelWithConfig()
		if err != nil {
			Logger.Error(err.Error())
		}
		baseRipple.Set(model)
	}

	Logger.Info("[gorm] initial DB finished")
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
		Format: "time=${time_rfc3339}, remote_ip=${remote_ip}, method=${method}, uri=${uri}, status=${status}\n",
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

// RegisterControllers register a controller for ripple App
func RegisterController(c Controller) {
	AddController(*baseRipple.Echo, c)
}

// RegisterModels registers models in the global ripple App.
func RegisterModels(models ...interface{}) {
	baseRipple.Model.AddModels(models...)
}

// Migrate runs migrations on the global ripple App.
func Migrate() {
	baseRipple.Model.AutoMigrateAll()
}

// Run run ripple application
func Run() {
	Migrate()
	Logger.Info("Ripple Run...")
	baseRipple.Echo.Run(standard.New(baseRipple.Config.Domain))
}
