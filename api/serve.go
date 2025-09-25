package api

import (
	"context"
	"net/http"
	"time"

	"github.com/go-profiler/grafana"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var e *echo.Echo

func Serve(addr string) error {
	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	// Grafana routes
	g := e.Group("/api/grafana")
	g.GET("", func(c echo.Context) error {
		return c.NoContent(200)
	})
	g.POST("/search", func(c echo.Context) error {
		return c.JSON(200, []string{"app1.cpu", "app1.mem", "app2.cpu", "app2.mem"})
	})
	g.POST("/query", func(c echo.Context) error {
		t := &grafana.Table{
			Columns: []grafana.Column{
				{Text: "time", Type: grafana.CollTime},
				{Text: "url", Type: grafana.CollUrl},
			},
			Rows: []grafana.Row{
				{time.Now(), "http://ololol.ru/heh"},
				{time.Now(), "http://ololol.ru/heh2"},
				{time.Now(), "http://'ololol.ru/heh3"},
			},
		}

		return c.JSON(200, t)
	})
	g.GET("/annotations", nil)
	g.GET("/tag-keys", nil)
	g.GET("/tag-values", nil)

	err := e.Start(addr)
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func Shutdown() {
	if e == nil {
		e.Logger.Error("server not started")
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	if err := e.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		e.Logger.Error(err)
	}
}
