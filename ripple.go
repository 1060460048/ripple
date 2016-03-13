package ripple

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var baseRipple *Ripple

// Ripple ripple struct
type Ripple struct {
	Echo   *echo.Echo
	Config *Config
	Model  *Model
}

func init() {
	baseRipple = NewRipple()

	// Initial the DB with "github.com/jinzhu/gorm"
	baseModel := baseRipple.Model
	if baseModel != nil && !baseModel.IsOpenDB() {
		err := baseModel.OpenWithConfig()
		if err != nil {
			baseRipple.Echo.Logger().Error(err.Error())
		}
	} else {
		model, err := NewModelWithConfig()
		if err != nil {
			baseRipple.Echo.Logger().Error(err.Error())
		}
		baseRipple.Set(model)
	}

	Logger().Info("[gorm] initial DB finished")
}

// NewRipple new a ripple instance
func NewRipple() *Ripple {
	newEcho := echo.New()

	r := &Ripple{}
	r.Set(newEcho)
	r.Set(NewConfig(newEcho))
	r.Set(NewModel())

	r.Echo.Use(mw.Logger())
	r.Echo.Use(mw.Recover())

	// Set render
	r.Echo.SetRenderer(NewRenderer())
	r.Echo.Static("/static", GetStaticPath())

	return r
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
	AddController(baseRipple.Echo, c)
}

// RegisterModels registers models in the global ripple App.
func RegisterModels(models  ...interface{}) {
	baseRipple.Model.AddModels(models...)
}

// Migrate runs migrations on the global ripple App.
func Migrate() {
	baseRipple.Model.AutoMigrateAll()
}

func Logger() echo.Logger{
	return baseRipple.Echo.Logger()
}

// Run run ripple application
func Run() {
	Migrate()
	baseRipple.Echo.Logger().Info("Ripple Run...")
	baseRipple.Echo.Run(baseRipple.Config.Domain)
}
