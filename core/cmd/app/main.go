package main

import (
	"fmt"
	"net/http"
	"os"

	"core/internal/pkg/redis"
	"github.com/labstack/echo/v4"
)

func configured() bool {
	_, ok := os.LookupEnv("NAME")
	if !ok {
		return ok
	}

	_, ok = os.LookupEnv("REDIS_HOST")
	if !ok {
		return ok
	}
	return ok
}

func main() {
	e := echo.New()

	e.GET("/liveness", func(c echo.Context) error {
		if ok := configured(); !ok {
			return c.String(http.StatusPreconditionFailed, "FAILED")
		}

		return c.String(http.StatusOK, "OK")
	})

	e.GET("/readiness", func(c echo.Context) error {
		if ok := configured(); !ok {
			return c.String(http.StatusPreconditionFailed, "FAILED")
		}

		redisHost := os.Getenv("REDIS_HOST")
		if err := redis.NewRedis(redisHost).Ping(); err != nil {
			return c.String(http.StatusFailedDependency, "FAILED")
		}

		return c.String(http.StatusOK, "OK")
	})

	e.GET("/", func(c echo.Context) error {
		if ok := configured(); !ok {
			return c.String(http.StatusPreconditionFailed, "FAILED")
		}

		name := os.Getenv("NAME")
		redisHost := os.Getenv("REDIS_HOST")
		return c.String(http.StatusOK, fmt.Sprintf("Hello %s, you did it! the redis host is: %s\n", name, redisHost))
	})

	e.Logger.Fatal(e.Start(":8080"))

}
