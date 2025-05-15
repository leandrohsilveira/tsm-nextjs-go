package setup

import "github.com/labstack/echo/v4"

type SetupRoutesFunc func(*echo.Echo)

type SetupRoutesResult struct {
	Err  error
	Func SetupRoutesFunc
}

func SetupRoutes(fn SetupRoutesFunc) SetupRoutesResult {
	return SetupRoutesResult{
		Err:  nil,
		Func: fn,
	}
}

func SetupRoutesError(err error) SetupRoutesResult {
	return SetupRoutesResult{
		Err:  err,
		Func: nil,
	}
}

func Routes(setups ...SetupRoutesResult) func(*echo.Echo) error {
	return func(app *echo.Echo) error {
		for _, setup := range setups {
			if setup.Err != nil {
				return setup.Err
			}
			setup.Func(app)
		}
		return nil
	}
}
