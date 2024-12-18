package main

import (
	"log"
	"log/slog"
)

func main() {
	config := config{
		addr: ":8080",
	}
	app := &application{
		config: config,
	}

	mux := app.mount()
	slog.Info("Starting API server", slog.String("addr", app.config.addr))
	log.Fatal(app.run(mux))
}
