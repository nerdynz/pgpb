//go:build pq

package core

import (
	"log/slog"
	"os"
	"strings"

	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
)

func connectDB(dbPath string) (*dbx.DB, error) {
	if strings.Contains(dbPath, "logs.db") {
		logdb := os.Getenv("LOGS_DATABASE")
		slog.Info("logdb", "logdb", logdb)
		return dbx.MustOpen("postgres", logdb)
	}
	db := os.Getenv("DATABASE")
	slog.Info("db", "db", db)
	return dbx.MustOpen("postgres", db)
}
