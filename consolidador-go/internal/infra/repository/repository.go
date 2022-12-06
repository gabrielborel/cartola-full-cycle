package repository

import (
	"database/sql"
	"github.com/gabrielborel/cartola-consolidador/internal/infra/db"
)

type Repository struct {
	dbConn *sql.DB
	*db.Queries
}
