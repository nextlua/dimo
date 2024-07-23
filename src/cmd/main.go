package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"log"
	"log/slog"
	"nextlua/dimo/internal/config"
	"nextlua/dimo/internal/handlers"
	"nextlua/dimo/internal/repository"
	"nextlua/dimo/internal/service"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func setupRouter() *gin.Engine {
	r := gin.New()

	/* sentry
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	*/

	envVariables, err := config.LoadEnvVars()
	if err != nil {
		log.Panic(err)
	}

	err = r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Printf("error:%s", err)
	}

	db := struct{}{}

	repo := repository.NewRepository(db)

	memoryCache := cache.New(5*time.Minute, 10*time.Minute)

	srv := service.NewService(
		slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})),
		repo,
		memoryCache,
		envVariables,
	)

	handlers.SetHandler(r, srv)

	r.Static("/swaggerui/", "docs/swagger-ui")

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	return r
}

func main() {

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "",

		TracesSampleRate: 1.0,
	}); err != nil {
		log.Panicf("error: %s", err)
	}

	gin.SetMode(gin.DebugMode)

	r := setupRouter()

	port := getPort()

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
