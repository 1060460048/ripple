package cache

import (
	"github.com/labstack/echo"
)

const rippleCacheStoreKey = "rippleCacheStore"

func Store(value interface{}) Cache {
	var cacher Cache
	switch v := value.(type) {
	case *echo.Context:
		cacher = v.Get(rippleCacheStoreKey).(Cache)
		if cacher == nil {
			panic("rippleStore not found, forget to Use Middleware ?")
		}
	default:

		panic("unknown Context")
	}

	if cacher == nil {
		panic("cache context not found")
	}

	return cacher
}

func rippleCacher(opt Options) echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			tagcache, err := New(opt)
			if err != nil {
				return err
			}

			c.Set(rippleCacheStoreKey, tagcache)

			if err = h(c); err != nil {
				return err
			}

			return nil
		}
	}
}
