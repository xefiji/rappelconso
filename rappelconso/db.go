package rappelconso

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // todo justify
)

func newDB(cfg database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Name,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
