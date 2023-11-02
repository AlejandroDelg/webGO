package db_repo

import (
	"database/sql"
	"github.com/AlejandroDelg/webgo/internal/config"
	"github.com/AlejandroDelg/webgo/internal/repository"
)

type postgresDBRepo struct{
	App *config.AppConfig
	DB *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig)repository.DbRepo {
	
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
	
}