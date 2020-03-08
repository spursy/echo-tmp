package main

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	num := 1
	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path: "/",
			MaxAge: 86400 * 7,
			HttpOnly: true,
		}
		num = num + 1
		sess.Values["foo"] = num
		sess.Save(c.Request(), c.Response())
		sessionValue := (sess.Values["foo"]).(int)

		return c.String(http.StatusOK,  strconv.Itoa(sessionValue))
	})

	e.Logger.Fatal(e.Start(":7001"))
}