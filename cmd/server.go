package main

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/xefiji/rappelconso/rappelconso"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("no env file loaded")
	}
}

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() error {
	reset := flag.Bool("reset", false, "reset all database taking all N records")
	flag.Parse()

	return rappelconso.Listen(
		rappelconso.WithResetMode(*reset),
		rappelconso.WithDB(
			env("DB_USER", "root"),
			env("DB_PASSWORD", "postgres"),
			env("DB_HOST", "localhost"),
			env("DB_PORT", "5432"),
			env("DB_NAME", "rappelconso"),
		),
	)
}

func env(name, fallback string) string {
	if val, ok := os.LookupEnv(name); ok {
		return val
	}
	return fallback
}
