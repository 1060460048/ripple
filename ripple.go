package ripple

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var baseRipple *Ripple

func init() {
	baseRipple = NewRipple()
}

type Ripple struct {
	Echo     *echo.Echo
	Config   *Config
	Renderer *Renderer
}

func NewRipple() *Ripple {
	newEcho := echo.New()
	LoadConfig(newEcho)

	r := &Ripple{
		Echo:     newEcho,
		Config:   GetConfig(),
		Renderer: NewRenderer(),
	}

	r.Echo.Use(mw.Logger())
	r.Echo.Use(mw.Recover())

	// Set render
	r.Echo.SetRenderer(r.Renderer)
	r.Echo.Static("/static", GetStaticPath())

	return r
}

func Run() {
	baseRipple.Echo.Logger().Debug("Ripple Run...")
	baseRipple.Echo.Run(baseRipple.Config.Domain)
}

func RegisterController(controller Controller) {
	GroupController(controller, baseRipple.Echo)
}
