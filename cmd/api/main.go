package main

import (
	"log"
	"log/slog"

	"github.com/wafiqpuyol/GO-Social/internal/env"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8001"),
	}
	app := &application{
		config: config,
	}

	mux := app.mount()
	slog.Info("Starting API server", slog.String("addr", app.config.addr))
	log.Fatal(app.run(mux))
}
