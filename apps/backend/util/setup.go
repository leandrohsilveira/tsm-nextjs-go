package util

import "github.com/labstack/echo/v4"

type SetupHook func(*echo.Echo)

type SetupRoutes struct {
	Err  error
	Hook SetupHook
}

func Setup(hook SetupHook) SetupRoutes {
	return SetupRoutes{
		Err:  nil,
		Hook: hook,
	}
}

func SetupError(err error) SetupRoutes {
	return SetupRoutes{
		Err:  err,
		Hook: nil,
	}
}
