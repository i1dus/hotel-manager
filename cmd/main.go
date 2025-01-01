package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"hotel-management/internal/handler"
	"hotel-management/internal/middleware"
	"log"
	"os"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	appCtx := context.Background()

	conn, err := pgx.Connect(appCtx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close(appCtx)

	botSettings := tele.Settings{
		Token:  os.Getenv("TELEGRAM_BOT_TOKEN"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(botSettings)
	if err != nil {
		log.Fatal(err)
		return
	}

	middleware := middleware.NewMiddleware(conn)
	b.Use(middleware.Logger())
	b.Use(middleware.PermissionCheck(appCtx))

	handlerController := handler.NewHandlerController(b, conn)
	handlerController.RegisterHandlers()

	b.Start()
}
