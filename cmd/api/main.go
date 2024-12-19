package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/wafiqpuyol/GO-Social/internal/db"
	"github.com/wafiqpuyol/GO-Social/internal/env"
	"github.com/wafiqpuyol/GO-Social/internal/store"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8001"),
		db: dbConfig{
			addr:               env.GetString("DB_ADDR", "postgres://admin:mySecretPassword@localhost/go_social?sslmode=disable"),
			maxOpenConnections: env.GetInt("MAX_OPEN_CONNECTIONS", 30),
			maxIdleConnections: env.GetInt("MAX_IDLE_CONNECTIONS", 30),
			maxIdleTime:        env.GetString("MAX_IDLE_TIME", "15m"),
		},
	}

	/* ------  initialize database connection ------ */
	slog.Info("Connecting to database", slog.String("addr", config.db.addr))
	fmt.Println("time ==>", config.db)
	db, err := db.NewDB(
		config.db.addr,
		config.db.maxIdleConnections,
		config.db.maxOpenConnections,
		config.db.maxIdleTime,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	store := store.NewStorage(db)
	slog.Info("DB connection established")

	/* ------ initialize server ------ */
	app := &application{
		config: config,
		store:  store,
	}
	mux := app.mount()
	slog.Info("Starting API server", slog.String("addr", app.config.addr))
	log.Fatal(app.run(mux))
}
