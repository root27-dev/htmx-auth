package main

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	v "github.com/root27-dev/htmx-auth/views"
	"net/http"
	"time"
)

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
func main() {

	users := make(map[string]string)

	e := echo.New()

	e.Use(middleware.Logger())

	// Pages

	e.GET("/", func(e echo.Context) error {

		return Render(e, 200, v.Index())

	})

	e.GET("/login", func(e echo.Context) error {

		return Render(e, 200, v.Login(false, "", false))

	})
	e.GET("/register", func(e echo.Context) error {

		return Render(e, 200, v.Register(false))

	})

	e.GET("/servelogin", func(e echo.Context) error {

		e.Response().Header().Set("Hx-Redirect", "/login")

		return nil
	})
	e.GET("/serveregister", func(e echo.Context) error {

		e.Response().Header().Set("Hx-Redirect", "/register")

		return nil
	})

	// Handlers

	e.POST("/api-register", func(e echo.Context) error {

		email := e.FormValue("email")
		password := e.FormValue("password")

		if _, ok := users[email]; ok {

			return Render(e, 200, v.Register(true))
		}

		users[email] = password

		e.Response().Header().Set("Hx-Redirect", "/login")

		return nil

	})

	e.POST("/api-login", func(e echo.Context) error {

		email := e.FormValue("email")

		password := e.FormValue("password")

		if p, ok := users[email]; ok {

			if p == password {
				// Demo for cookie example
				cookie := new(http.Cookie)
				cookie.Name = "example"
				cookie.Value = email
				cookie.Expires = time.Now().Add(24 * time.Hour)
				e.SetCookie(cookie)

				return Render(e, 200, v.Login(true, email, false))
			}
		}

		return Render(e, 200, v.Login(false, "", true))

	})

	e.Logger.Fatal(e.Start(":8080"))

}
